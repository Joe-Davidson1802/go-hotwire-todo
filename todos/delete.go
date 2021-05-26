package todos

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/joe-davidson1802/go-hotwire-todo/store"
)

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
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

	s := store.NewTodoStore("", "Todo")

	if err := s.DeleteTodo(r.Context(), int64(id)); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
