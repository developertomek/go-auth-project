package main

import (
	"fmt"
	"net/http"

	"github.com/developertomek/go-auth-project/api"
	"github.com/developertomek/go-auth-project/db"
	"github.com/developertomek/go-auth-project/routes"
)

func main() {
	fmt.Println("Program starting...")

	database, err := db.Open()
	if err != nil {
		panic(err)
	}

	userStore := db.NewSQLiteUserStore(database)
	userHandler := api.NewUserHandler(userStore)

	r := routes.SetupRoutes(*userHandler)

	http.ListenAndServe(":8080", r)

}
