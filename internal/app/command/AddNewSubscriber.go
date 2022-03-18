package command

import (
	"context"

	"github.com/sirupsen/logrus"
)

type AddNewSubscriberHandler struct {
	writeToModel AddNewSubscriberReadModel
}

type AddNewSubscriberReadModel interface {
	UpdatedToken(ctx context.Context, token string) error
}


func NewAddNewSubscriberHandler(tokenRepo AddNewSubscriberReadModel) AddNewSubscriberHandler {
	if tokenRepo == nil {
		panic("nil tokenRepo")
	}

	return AddNewSubscriberHandler{
		writeToModel: tokenRepo,
	}
}

func (c AddNewSubscriberHandler) Handle(ctx context.Context, query string) (err error) {
	defer func() {
		logrus.WithError(err).Debug("AddNewSubscriberHandler Executed")
	}()

	return c.writeToModel.UpdatedToken(ctx, query)
}