package todos

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/joe-davidson1802/go-hotwire-todo/store"
	"github.com/joe-davidson1802/go-hotwire-todo/views"
)

func CompleteTodoHandler(w http.ResponseWriter, r *http.Request) {
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

	s := store.NewTodoStore("", "Todo")

	t, err := s.CompleteTodo(r.Context(), int64(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/vnd.turbo-stream.html")

	err = views.TodoRow(*t, "replace", strconv.Itoa(int(t.ID.ID))).Render(r.Context(), w)
}
