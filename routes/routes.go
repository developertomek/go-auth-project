package routes

import (
	"fmt"
	"net/http"

	"github.com/developertomek/go-auth-project/api"
	"github.com/developertomek/go-auth-project/middleware"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(userHandler api.UserHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", api.HandleHelloWorld)
	r.Post("/echo", api.HandleEchoUser)
	r.Post("/register", userHandler.HandlerRegisterUser)
	r.Post("/login", userHandler.HandlerLoginUser)

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.UserAuthentication)
		r.Get("/test", HandleTest)
	})
	return r
}

func HandleTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "test working")
}
