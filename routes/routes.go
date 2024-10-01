package routes

import (
	"github.com/developertomek/go-auth-project/api"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(userHandler api.UserHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", api.HandleHelloWorld)
	r.Post("/echo", api.HandleEchoUser)
	r.Post("/register", userHandler.HandlerRegisterUser)
	r.Post("/login", userHandler.HandlerLoginUser)
	return r
}
