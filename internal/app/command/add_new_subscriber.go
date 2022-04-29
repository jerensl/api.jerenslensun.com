package command

import (
	"context"

	"github.com/jerensl/api.jerenslensun.com/internal/logs/errors"
)

type AddNewSubscriberHandler struct {
	writeToModel AddNewSubscriberWriteModel
}

type AddNewSubscriberWriteModel interface {
	UpdatedToken(token string) error
}


func NewAddNewSubscriberHandler(tokenRepo AddNewSubscriberWriteModel) AddNewSubscriberHandler {
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