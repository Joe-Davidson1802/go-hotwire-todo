package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/joe-davidson1802/go-hotwire-todo/models"
	"github.com/joe-davidson1802/go-hotwire-todo/views"
)

var decoder = schema.NewDecoder()

func postTodoHandler(w http.ResponseWriter, r *http.Request) {
	var t models.Todo

	fmt.Println("Received POST to /create-todo")

	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = decoder.Decode(&t, r.PostForm)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s := NewTodoStore("plasma-myth-310415", "Todo")

	if err = s.PostTodo(r.Context(), &t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	payload, err := t.MarshalJSON()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received DELETE to /delete-todo")

	idParam := r.URL.Query().Get("id")

	if idParam == "" {
		http.Error(w, "id was missing", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s := NewTodoStore("plasma-myth-310415", "Todo")

	if err := s.DeleteTodo(r.Context(), int64(id)); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func getTodoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received GET to /get-todo")

	idParam := r.URL.Query().Get("id")

	if idParam == "" {
		http.Error(w, "id was missing", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s := NewTodoStore("plasma-myth-310415", "Todo")

	t, err := s.GetTodo(r.Context(), int64(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	payload, err := t.MarshalJSON()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func getTodosHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received GET to /get-todos")

	maxParam := r.URL.Query().Get("max")

	if maxParam == "" {
		http.Error(w, "id was missing", http.StatusBadRequest)
		return
	}

	max, err := strconv.Atoi(maxParam)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s := NewTodoStore("plasma-myth-310415", "Todo")

	ts, err := s.GetTodos(r.Context(), max)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	err = views.ListView(*ts).Render(r.Context(), w)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received PUT to /complete-todo")

	idParam := r.URL.Query().Get("id")

	if idParam == "" {
		http.Error(w, "id was missing", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s := NewTodoStore("plasma-myth-310415", "Todo")

	t, err := s.CompleteTodo(r.Context(), int64(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	payload, err := t.MarshalJSON()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	err := views.RenderHome().Render(r.Context(), w)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/create-todo", postTodoHandler).Methods("POST")
	r.HandleFunc("/delete-todo", deleteTodoHandler).Methods("DELETE")
	r.HandleFunc("/get-todo", getTodoHandler).Methods("GET")
	r.HandleFunc("/get-todos", getTodosHandler).Methods("GET")
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/complete-todo", completeTodoHandler).Methods("PUT")

	err := http.ListenAndServe(":80", r)
	log.Fatal(err)
}
