package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type User struct {
	Name string `json: "name"`
}

var userChahe = make(map[int]User)

var chacheMutex sync.RWMutex

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

	chacheMutex.Lock()
	userChahe[len(userChahe)+1] = user
	chacheMutex.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

func getUser(
	w http.ResponseWriter,
	r *http.Request,
) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}
	chacheMutex.RLock()
	user, ok := userChahe[id]
	chacheMutex.RUnlock()

	if !ok {
		http.Error(
			w,
			"User not found",
			http.StatusNotFound,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(user)
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func deleteUser(
	w http.ResponseWriter,
	r *http.Request,
) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	if _, ok := userChahe[id]; !ok {
		http.Error(
			w,
			"User not found",
			http.StatusNotFound,
		)
		return
	}
	chacheMutex.Lock()
	delete(userChahe, id)
	chacheMutex.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)

	mux.HandleFunc("POST /users", createUser)
	mux.HandleFunc("GET /users/{id}", getUser)
	mux.HandleFunc("DELETE /users/{id}", deleteUser)

	fmt.Println("Server is on 8080")
	http.ListenAndServe(":8080", mux)
}
