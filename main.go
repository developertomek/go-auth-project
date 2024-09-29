package main

import (
	"fmt"
	"net/http"

	"github.com/developertomek/go-auth-project/routes"
)

func main() {
	fmt.Println("Program starting...")
	r := routes.SetupRoutes()
	http.ListenAndServe(":8080", r)
}
