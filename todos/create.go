package todos

import (
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/joe-davidson1802/go-hotwire-todo/models"
	"github.com/joe-davidson1802/go-hotwire-todo/store"
	"github.com/joe-davidson1802/go-hotwire-todo/views"
)

var decoder = schema.NewDecoder()

func PostTodoHandler(w http.ResponseWriter, r *http.Request) {
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

	s := store.NewTodoStore("", "Todo")

	if err = s.PostTodo(r.Context(), &t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/vnd.turbo-stream.html")

	err = views.TodoRow(t, "append", "todo_lister").Render(r.Context(), w)
}
