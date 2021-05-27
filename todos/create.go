package todos

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/joe-davidson1802/go-hotwire-todo/models"
	"github.com/joe-davidson1802/go-hotwire-todo/store"
	"github.com/joe-davidson1802/go-hotwire-todo/views"
)

var decoder = schema.NewDecoder()

type CreateHandler struct{}

func (h CreateHandler) CanHandleModel(m string) bool {
	return m == models.Todo{}.ModelName()
}

func (h CreateHandler) HandleRequest(w http.ResponseWriter, r *http.Request) (error, models.Model) {
	var t models.Todo

	fmt.Println("Received POST to /create-todo")

	err := r.ParseForm()

	if err != nil {
		return err, nil
	}

	err = decoder.Decode(&t, r.PostForm)

	if err != nil {
		return err, nil
	}

	s := store.NewTodoStore("", "Todo")

	if err = s.PostTodo(r.Context(), &t); err != nil {
		return err, nil
	}

	return nil, t
}

func (h CreateHandler) RenderPage(ctx context.Context, m models.Model, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/html")

	t := m.(*models.Todo)

	v := views.Layout(t.Title, views.TodoRow(*t, "append", "todo_lister", true))

	err := v.Render(ctx, w)

	return err
}

func (h CreateHandler) RenderStream(ctx context.Context, m models.Model, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/vnd.turbo-stream.html")

	t := m.(*models.Todo)

	err := views.TodoRow(*t, "append", "todo_lister", true).Render(ctx, w)

	return err
}
