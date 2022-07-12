package query

import (
	"context"

	"github.com/jerensl/api.jerenslensun.com/internal/domain/notification"
	"github.com/jerensl/api.jerenslensun.com/internal/logs/errors"
)

type GetStatusSubscriberHandler struct {
	tokenRepo notification.Repository
}


func NewGetStatusSubscriberHandler(tokenRepo notification.Repository) GetStatusSubscriberHandler {
	if tokenRepo == nil {
		panic("nil tokenRepo")
	}

	return GetStatusSubscriberHandler{
		tokenRepo: tokenRepo,
	}
}

func (c GetStatusSubscriberHandler) Handle(ctx context.Context, query string) (*notification.Token, error) {
	token, _, err := c.tokenRepo.GetToken(query)
	if err != nil {
		return token, errors.NewSlugError(err.Error(), "unable to get token")
	}

	return token, nil
}