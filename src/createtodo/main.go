package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func putTodoHandler(w http.ResponseWriter, r *http.Request) {
	var t Todo

	fmt.Println("Received POST to /create-todo")

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte(t.Title))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/create-todo", putTodoHandler).Methods("POST")

	err := http.ListenAndServe(":80", r)
	log.Fatal(err)
}
