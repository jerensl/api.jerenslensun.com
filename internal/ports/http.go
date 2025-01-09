package ports

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-chi/render"
	"github.com/jerensl/api.jerenslensun.com/internal/app"
	"github.com/jerensl/api.jerenslensun.com/internal/app/command"
	"github.com/jerensl/api.jerenslensun.com/internal/logs/httperr"
)

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=../../api/openapi/types.cfg.yaml ../../api/openapi/notification.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=../../api/openapi/server.cfg.yaml ../../api/openapi/notification.yaml

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

	cmd := command.AddNewSubscriber{
		TokenID: newSubscriber.TokenID,
		UpdateAt: newSubscriber.UpdatedAt,
	}

	err := h.app.Commands.AddNewSubscriber.Handle(r.Context(), cmd)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}
	
	render.Status(r, http.StatusNoContent)
}

func (h HttpServer) SubscriberStatus(w http.ResponseWriter, r *http.Request) {
	var subscriber Subscriber
	if err := json.NewDecoder(r.Body).Decode(&subscriber); err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	token, err := h.app.Queries.GetStatusSubscriber.Handle(r.Context(), subscriber.TokenID)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	status := Status{
		IsActive: token.IsActive(),
		UpdatedAt: token.UpdatedAt(),
	}

	render.Respond(w, r, status)
}

func (h HttpServer) UnsubscribeNotification(w http.ResponseWriter, r *http.Request) {
	var subscriber Subscriber
	if err := json.NewDecoder(r.Body).Decode(&subscriber); err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	cmd := command.Unsubscribe{
		TokenID: subscriber.TokenID,
		UpdateAt: subscriber.UpdatedAt,
	}

	err := h.app.Commands.Unsubscribe.Handle(r.Context(), cmd)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	render.Status(r, http.StatusNoContent)
}

func (h HttpServer) SendNotification(w http.ResponseWriter, r *http.Request) {
	var message Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	token := r.Header.Get("X-API-KEY")
	if token != os.Getenv("API_KEY") {
		httperr.Unauthorised("invalid token", nil,w, r)
		return
	}

	subscriber, err := h.app.Queries.GetAllSubscriber.Handle(r.Context())
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	h.app.Commands.SendNotification.Handle(r.Context(), subscriber, message.Title, message.Message)

	w.WriteHeader(http.StatusOK)
}

func (h HttpServer) SubscriberStats(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-API-KEY")
	if token != os.Getenv("API_KEY") {
		httperr.Unauthorised("invalid token", nil,w, r)
		return
	}

	subsStats, err := h.app.Queries.GetStatsSubscriber.Handle()
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	stats := Stats{
		TotalSubs: subsStats.TotalSubs(),
		TotalActiveSubs: subsStats.TotalActiveSubs(),
		TotalInactiveSubs: subsStats.TotalInactiveSubs(),
	}

	render.Respond(w, r, stats)
}