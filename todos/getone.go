package todos

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/joe-davidson1802/go-hotwire-todo/store"
)

func GetTodoHandler(w http.ResponseWriter, r *http.Request) {
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

	s := store.NewTodoStore("", "Todo")

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
