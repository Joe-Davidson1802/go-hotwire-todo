package todos

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/joe-davidson1802/go-hotwire-todo/store"
	"github.com/joe-davidson1802/go-hotwire-todo/views"
)

func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
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

	s := store.NewTodoStore("", "Todo")

	ts, err := s.GetTodos(r.Context(), max)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	err = views.Layout("List Todos", views.ListView(*ts)).Render(r.Context(), w)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
