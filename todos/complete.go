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

type CompleteHandler struct{}

func (h CompleteHandler) CanHandleModel(m string) bool {
	return m == models.Todo{}.ModelName()
}

func (h CompleteHandler) HandleRequest(w http.ResponseWriter, r *http.Request) (error, models.Model) {
	fmt.Println("Received PUT to /complete-todo")

	idParam := r.URL.Query().Get("id")

	if idParam == "" {
		return errors.New("id was missing"), nil
	}

	id, err := strconv.Atoi(idParam)

	if err != nil {
		return err, nil
	}

	s := store.NewTodoStore("", "Todo")

	t, err := s.CompleteTodo(r.Context(), int64(id))
	if err != nil {
		return err, nil
	}

	return nil, t
}

func (h CompleteHandler) RenderPage(ctx context.Context, m models.Model, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/html")

	t := m.(models.Todo)

	v := views.Layout(t.Title, views.TodoRow(t, "replace", "row_"+strconv.Itoa(int(t.ID.ID)), true))

	err := v.Render(ctx, w)

	return err
}

func (h CompleteHandler) RenderStream(ctx context.Context, m models.Model, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/vnd.turbo-stream.html")

	t := m.(models.Todo)

	err := views.TodoRow(t, "replace", "row_"+strconv.Itoa(int(t.ID.ID)), true).Render(ctx, w)

	return err
}
