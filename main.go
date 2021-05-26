package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	r.HandleFunc("/create-todo", todos.PostTodoHandler).Methods("POST")
	r.HandleFunc("/delete-todo", todos.DeleteTodoHandler).Methods("DELETE")
	r.HandleFunc("/get-todo", todos.GetTodoHandler).Methods("GET")
	r.HandleFunc("/get-todos", todos.GetTodosHandler).Methods("GET")
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/complete-todo", todos.CompleteTodoHandler).Methods("PUT")

	err := http.ListenAndServe(":80", r)
	log.Fatal(err)
}
