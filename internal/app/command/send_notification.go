package command

import (
	"github.com/jerensl/jerens-web-api/internal/logs/errors"
)

type SendNotificationHandler struct {
	writeToModel SendNotificationReadModel
}

type SendNotificationReadModel interface {
	SendNotification(token []string, title string, message string) error
}


func NewSendNotificationHandler(tokenRepo SendNotificationReadModel) SendNotificationHandler {
	if tokenRepo == nil {
		panic("nil tokenRepo")
	}

	return SendNotificationHandler{
		writeToModel: tokenRepo,
	}
}

func (c SendNotificationHandler) Handle(listOfToken []string, title string, message string) (error) {
	err := c.writeToModel.SendNotification(listOfToken, title, message)
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable to send notifications")
	}

	return nil
}