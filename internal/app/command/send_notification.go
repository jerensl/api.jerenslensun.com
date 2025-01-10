package command

import (
	"context"

	"github.com/jerensl/api.jerenslensun.com/internal/decorator"
	"github.com/jerensl/api.jerenslensun.com/internal/logs/errors"
	"github.com/sirupsen/logrus"
)

type SendNotification struct {
	ListOfToken []string 
	Title 		string
	Message 	string
}

type SendNotificationHandler decorator.CommandHandler[SendNotification]

type sendNotificationHandler struct {
	notificationService NotificationService
}

func NewSendNotificationHandler(notificationService NotificationService, logger *logrus.Entry, metricsClient decorator.MetricsClient) SendNotificationHandler {
	if notificationService == nil {
		panic("nil notificationService")
	}

	return decorator.ApplyCommandDecorator[SendNotification](
		sendNotificationHandler{notificationService: notificationService},
		logger,
		metricsClient,
	)
}

func (c sendNotificationHandler) Handle(ctx context.Context, cmd SendNotification) (error) {
	err := c.notificationService.SendNotification(ctx, cmd.ListOfToken, cmd.Title, cmd.Message)
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable to send notifications")
	}

	return nil
}