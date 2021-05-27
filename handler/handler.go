package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/joe-davidson1802/go-hotwire-todo/models"
)

type RequestHandler interface {
	HandleRequest(http.ResponseWriter, *http.Request) (error, models.Model)
	RenderPage(context.Context, models.Model, http.ResponseWriter) error
	RenderStream(context.Context, models.Model, http.ResponseWriter) error
	CanHandleModel(string) bool
}

type TurboHandler struct {
	handler RequestHandler
}

func New(h RequestHandler) TurboHandler {
	return TurboHandler{handler: h}
}

func (h TurboHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err, m := h.handler.HandleRequest(w, r)

	if !h.handler.CanHandleModel(m.ModelName()) {
		err = errors.New("Handler cannot handler model of type " + m.ModelName())
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ct := r.Header.Get("Accept")

	if ct == "text/vnd.turbo-stream.html" {
		err = h.handler.RenderStream(r.Context(), m, w)
	} else {
		err = h.handler.RenderPage(r.Context(), m, w)
	}
}
