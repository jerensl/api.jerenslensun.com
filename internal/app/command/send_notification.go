package command

import (
	"context"

	"github.com/jerensl/api.jerenslensun.com/internal/logs/errors"
)

type SendNotificationHandler struct {
	notificationService NotificationService
}

func NewSendNotificationHandler(notificationService NotificationService) SendNotificationHandler {
	if notificationService == nil {
		panic("nil notificationService")
	}

	return SendNotificationHandler{
		notificationService: notificationService,
	}
}

func (c SendNotificationHandler) Handle(ctx context.Context, listOfToken []string, title string, message string) (error) {
	err := c.notificationService.SendNotification(ctx, listOfToken, title, message)
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable to send notifications")
	}

	return nil
}