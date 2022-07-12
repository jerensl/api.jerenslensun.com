package command

import (
	"context"

	"github.com/jerensl/api.jerenslensun.com/internal/domain/notification"
	"github.com/jerensl/api.jerenslensun.com/internal/logs/errors"
)

type AddNewSubscriberHandler struct {
	tokenRepo notification.Repository
}

func NewAddNewSubscriberHandler(tokenRepo notification.Repository) AddNewSubscriberHandler {
	if tokenRepo == nil {
		panic("nil tokenRepo")
	}

	return AddNewSubscriberHandler{
		tokenRepo: tokenRepo,
	}
}

func (c AddNewSubscriberHandler) Handle(ctx context.Context, tokenID string, updateAt int64) (*notification.Token, error) {
	token, err := c.tokenRepo.UpdatedToken(tokenID, updateAt)
	if err != nil {
		return nil, errors.NewSlugError(err.Error(), "unable to add subscriber")
	}

	return token, nil
}