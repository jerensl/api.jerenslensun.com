package command

import (
	"context"

	"github.com/jerensl/api.jerenslensun.com/internal/logs/errors"
)

type SendNotificationHandler struct {
	writeToModel SendNotificationReadModel
}

type SendNotificationReadModel interface {
	SendNotification(ctx context.Context, token []string, title string, message string) error
}


func NewSendNotificationHandler(tokenRepo SendNotificationReadModel) SendNotificationHandler {
	if tokenRepo == nil {
		panic("nil tokenRepo")
	}

	return SendNotificationHandler{
		writeToModel: tokenRepo,
	}
}

func (c SendNotificationHandler) Handle(ctx context.Context, listOfToken []string, title string, message string) (error) {
	err := c.writeToModel.SendNotification(ctx, listOfToken, title, message)
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable to send notifications")
	}

	return nil
}