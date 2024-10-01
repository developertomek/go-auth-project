package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/developertomek/go-auth-project/types"
)

type UserStore interface {
	CreateUser(context.Context, *types.User) (*types.User, error)
	GetUserByEmail(context.Context, string) (*types.User, error)
}

type SQLiteUserStore struct {
	db *sql.DB
}

func NewSQLiteUserStore(db *sql.DB) *SQLiteUserStore {
	return &SQLiteUserStore{
		db: db,
	}
}

func (u *SQLiteUserStore) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	query := `
    INSERT INTO users (email, password_hash)
    VALUES (?,?)
    RETURNING id;`

	var userID string
	err := u.db.QueryRowContext(ctx, query, user.Email, user.PasswordHash).Scan(&userID)

	if err != nil {
		return nil, fmt.Errorf("createUser: %w", err)
	}

	user.ID = userID

	return user, nil
}

func (u *SQLiteUserStore) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	query := `SELECT id, password_hash, email FROM users WHERE email = ?`
	var user types.User
	err := u.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.PasswordHash, &user.Email)

	if err != nil {
		return nil, fmt.Errorf("getUserByEmail: %w", err)
	}

	return &user, nil
}
