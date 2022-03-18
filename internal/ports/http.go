package ports

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jerensl/jerens-web-api/internal/app"
	"github.com/jerensl/jerens-web-api/internal/logs/httperr"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(application app.Application) HttpServer {
	return HttpServer{
		app: application,
	}
}

func (h HttpServer) Subscribe(w http.ResponseWriter, r *http.Request) {
	var newSubscriber NewSubscriber
	if err := json.NewDecoder(r.Body).Decode(&newSubscriber); err != nil {
		return
	}

	ctx := context.Background()

	err := h.app.Commands.AddNewSubscriber.Handle(ctx, newSubscriber.Token)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusCreated)
}