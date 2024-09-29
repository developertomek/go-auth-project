package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/developertomek/go-auth-project/types"
)

func HandleHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func HandleEchoUser(w http.ResponseWriter, r *http.Request) {
	var user types.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "this is the name: %s", user.Name)
}
