package command

import (
	"context"

	"github.com/jerensl/jerens-web-api/internal/logs/errors"
)

type AddNewSubscriberHandler struct {
	writeToModel AddNewSubscriberReadModel
}

type AddNewSubscriberReadModel interface {
	UpdatedToken(token string) error
}


func NewAddNewSubscriberHandler(tokenRepo AddNewSubscriberReadModel) AddNewSubscriberHandler {
	if tokenRepo == nil {
		panic("nil tokenRepo")
	}

	return AddNewSubscriberHandler{
		writeToModel: tokenRepo,
	}
}

func (c AddNewSubscriberHandler) Handle(ctx context.Context, query string) (error) {
	err := c.writeToModel.UpdatedToken(query)
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable to add subscriber")
	}

	return nil
}