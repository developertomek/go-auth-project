package types

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

type CreateUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(params CreateUser) (*User, error) {
	hpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), 10)
	if err != nil {
		return nil, err
	}

	return &User{
		Email:        params.Email,
		PasswordHash: string(hpw),
	}, nil
}

func ValidatePassword(hashPassword string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password)) == nil
}

func CreateToken(user User) string {
	now := time.Now()
	validUntil := now.Add(time.Hour * 4).Unix()

	claims := jwt.MapClaims{
		"id":      user.ID,
		"email":   user.Email,
		"expires": validUntil,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims, nil)
	secret := "TEST_SECRET"

	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Printf("failed to sign token: %w", err)
	}

	return tokenStr
}
