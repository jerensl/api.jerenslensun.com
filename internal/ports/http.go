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

func (h HttpServer) SubscribeNotification(w http.ResponseWriter, r *http.Request) {
	var newSubscriber Subscriber
	if err := json.NewDecoder(r.Body).Decode(&newSubscriber); err != nil {
		httperr.RespondWithSlugError(err, w, r)
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

func (h HttpServer) SendNotification(w http.ResponseWriter, r *http.Request) {
	var message Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	ctx := context.Background()

	subscriber, err := h.app.Queries.GetAllSubscriber.Handle(ctx)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	h.app.Commands.SendNotification.Handle(subscriber, message.Title, message.Message)

	w.WriteHeader(http.StatusOK)
}

func (h HttpServer) SubscriberStatus(w http.ResponseWriter, r *http.Request) {
	var subscriber Subscriber
	if err := json.NewDecoder(r.Body).Decode(&subscriber); err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	ctx := context.Background()

	isSubscriber, err := h.app.Queries.GetStatusSubscriber.Handle(ctx, subscriber.Token)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	status := Status{
		Status: isSubscriber,
	}
	if err := json.NewEncoder(w).Encode(status); err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h HttpServer) UnsubscribeNotification(w http.ResponseWriter, r *http.Request) {
	var subscriber Subscriber
	if err := json.NewDecoder(r.Body).Decode(&subscriber); err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	err := h.app.Commands.Unsubscribe.Handle(subscriber.Token)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
}