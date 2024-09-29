package routes

import (
	"github.com/developertomek/go-auth-project/api"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", api.HandleHelloWorld)
	r.Post("/echo", api.HandleEchoUser)
	return r
}
