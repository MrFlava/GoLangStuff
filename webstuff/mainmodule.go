package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json: "name"`
}

var userChahe = make(map[int]User)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func createUser(
	w http.ResponseWriter,
	r *http.Request,
) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Name == "" {
		http.Error(
			w,
			"name is reqired",
			http.StatusBadRequest,
		)
		return
	}

	userChahe[len(userChahe)+1] = user

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)

	mux.HandleFunc("POST /users", createUser)

	fmt.Println("Server is on 8080")
	http.ListenAndServe(":8080", mux)
}
