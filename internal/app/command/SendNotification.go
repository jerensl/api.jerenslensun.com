package command

import (
	"context"

	"github.com/sirupsen/logrus"
)

type SendNotificationHandler struct {
	writeToModel SendNotificationReadModel
}

type SendNotificationReadModel interface {
	SendNotification(token []string, messageClient string) error
}


func NewSendNotificationHandler(tokenRepo SendNotificationReadModel) SendNotificationHandler {
	if tokenRepo == nil {
		panic("nil tokenRepo")
	}

	return SendNotificationHandler{
		writeToModel: tokenRepo,
	}
}

func (c SendNotificationHandler) Handle(ctx context.Context, listOfToken []string, message string) (err error) {
	defer func() {
		logrus.WithError(err).Debug("SendNotificationHandler Executed")
	}()

	return c.writeToModel.SendNotification(listOfToken, message)
}