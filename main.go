package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joe-davidson1802/go-hotwire-todo/handler"
	"github.com/joe-davidson1802/go-hotwire-todo/todos"
	"github.com/joe-davidson1802/go-hotwire-todo/views"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	err := views.Layout("Home", views.RenderHome()).Render(r.Context(), w)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func main() {
	r := mux.NewRouter()
	r.Handle("/create-todo", handler.New(todos.CreateHandler{})).Methods("POST")
	r.Handle("/delete-todo", handler.New(todos.DeleteHandler{})).Methods("DELETE")
	r.HandleFunc("/get-todo", todos.GetTodoHandler).Methods("GET")
	r.Handle("/get-todos", handler.New(todos.GetAllHandler{})).Methods("GET")
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.Handle("/complete-todo", handler.New(todos.CompleteHandler{})).Methods("PUT")

	err := http.ListenAndServe(":80", r)
	log.Fatal(err)
}
