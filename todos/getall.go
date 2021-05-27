package todos

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/joe-davidson1802/go-hotwire-todo/models"
	"github.com/joe-davidson1802/go-hotwire-todo/store"
	"github.com/joe-davidson1802/go-hotwire-todo/views"
)

type GetAllHandler struct{}

func (h GetAllHandler) CanHandleModel(m string) bool {
	return m == models.Todo{}.ModelName()
}

func (h GetAllHandler) HandleRequest(w http.ResponseWriter, r *http.Request) (error, models.Model) {
	fmt.Println("Received GET to /get-todos")

	maxParam := r.URL.Query().Get("max")

	if maxParam == "" {
		return errors.New("id was missing"), nil
	}

	max, err := strconv.Atoi(maxParam)

	if err != nil {
		return err, nil
	}

	s := store.NewTodoStore("", "Todo")

	ts, err := s.GetTodos(r.Context(), max)
	if err != nil {
		return err, nil
	}

	return nil, models.TodoList{Todos: ts}
}

func (h GetAllHandler) RenderPage(ctx context.Context, m models.Model, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/html")

	t := m.(models.TodoList)

	err := views.Layout("List Todos", views.ListView(*t.Todos)).Render(ctx, w)

	return err
}

func (h GetAllHandler) RenderStream(ctx context.Context, m models.Model, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/vnd.turbo-stream.html")

	t := m.(models.TodoList)

	err := views.ListView(*t.Todos).Render(ctx, w)

	return err
}
