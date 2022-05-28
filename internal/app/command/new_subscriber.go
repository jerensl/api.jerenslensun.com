package command

import (
	"context"

	"github.com/jerensl/api.jerenslensun.com/internal/domain"
	"github.com/jerensl/api.jerenslensun.com/internal/logs/errors"
)

type AddNewSubscriberHandler struct {
	tokenRepo domain.Repository
}

func NewAddNewSubscriberHandler(tokenRepo domain.Repository) AddNewSubscriberHandler {
	if tokenRepo == nil {
		panic("nil tokenRepo")
	}

	return AddNewSubscriberHandler{
		tokenRepo: tokenRepo,
	}
}

func (c AddNewSubscriberHandler) Handle(ctx context.Context, query string) (error) {
	err := c.tokenRepo.UpdatedToken(query)
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable to add subscriber")
	}

	return nil
}