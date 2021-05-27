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

type DeleteHandler struct{}

func (h DeleteHandler) CanHandleModel(m string) bool {
	return m == models.Todo{}.ModelName()
}

func (h DeleteHandler) HandleRequest(w http.ResponseWriter, r *http.Request) (error, models.Model) {
	fmt.Println("Received DELETE to /delete-todo")

	idParam := r.URL.Query().Get("id")

	if idParam == "" {
		return errors.New("id is missing"), nil
	}

	id, err := strconv.Atoi(idParam)

	if err != nil {
		return err, nil
	}

	s := store.NewTodoStore("", "Todo")

	if err := s.DeleteTodo(r.Context(), int64(id)); err != nil {
		return err, nil
	}

	return nil, models.TodoId{Value: idParam}
}

func (h DeleteHandler) RenderPage(ctx context.Context, m models.Model, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/html")
	return errors.New("Not implemented")
}

func (h DeleteHandler) RenderStream(ctx context.Context, m models.Model, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/vnd.turbo-stream.html")

	t := m.(models.TodoId)

	err := views.RemoveTodoRow(t.Value, true).Render(ctx, w)

	return err
}
