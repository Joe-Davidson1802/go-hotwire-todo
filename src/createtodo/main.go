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

	s := NewTodoStore("plasma-myth-310415", "Todo")

	newT, err := s.PutTodo(r.Context(), t)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte(newT.Title))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/create-todo", putTodoHandler).Methods("POST")

	err := http.ListenAndServe(":80", r)
	log.Fatal(err)
}
