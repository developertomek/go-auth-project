package api

import (
	"encoding/json"
	"net/http"

	"github.com/developertomek/go-auth-project/db"
	"github.com/developertomek/go-auth-project/types"
)

type UserHandler struct {
	userStore db.UserStore
}

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User  *types.User `json:"user"`
	Token string      `json:"token"`
}

func NewUserHandler(us db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: us,
	}
}

func (u *UserHandler) HandlerRegisterUser(w http.ResponseWriter, r *http.Request) {
	var params types.CreateUser
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	user, err := types.NewUser(params)
	if err != nil {
		http.Error(w, "invalid new user", http.StatusBadRequest)
		return
	}

	user, err = u.userStore.CreateUser(ctx, user)
	if err != nil {
		http.Error(w, "invalid new user", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "could not register new user", http.StatusInternalServerError)
	}
}

func (u *UserHandler) HandlerLoginUser(w http.ResponseWriter, r *http.Request) {
	var loginParams LoginParams
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&loginParams); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	user, err := u.userStore.GetUserByEmail(ctx, loginParams.Email)
	if err != nil {
		http.Error(w, "invalid login user", http.StatusBadRequest)
		return
	}

	isValid := types.ValidatePassword(user.PasswordHash, loginParams.Password)
	if isValid == false {
		http.Error(w, "invalid email or password", http.StatusBadRequest)
		return
	}

	response := LoginResponse{
		User:  user,
		Token: types.CreateToken(*user),
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "could not login user", http.StatusInternalServerError)
	}
}
