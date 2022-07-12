package query

import (
	"context"

	"github.com/jerensl/api.jerenslensun.com/internal/domain/notification"
	"github.com/jerensl/api.jerenslensun.com/internal/logs/errors"
)

type GetAllSubscriberHandler struct {
	tokenRepo notification.Repository
}


func NewGetAllSubscriberHandler(tokenRepo notification.Repository) GetAllSubscriberHandler {
	if tokenRepo == nil {
		panic("nil tokenRepo")
	}

	return GetAllSubscriberHandler{
		tokenRepo: tokenRepo,
	}
}

func (c GetAllSubscriberHandler) Handle(ctx context.Context) ([]string, error) {
	subscriber, err := c.tokenRepo.GetAllToken()

	if err != nil {
		return nil, errors.NewSlugError(err.Error(), "unable to get all token")
	}

	return subscriber, nil
}