package query

import (
	"context"

	"github.com/jerensl/api.jerenslensun.com/internal/domain"
	"github.com/jerensl/api.jerenslensun.com/internal/logs/errors"
)

type GetStatusSubscriberHandler struct {
	tokenRepo domain.Repository
}


func NewGetStatusSubscriberHandler(tokenRepo domain.Repository) GetStatusSubscriberHandler {
	if tokenRepo == nil {
		panic("nil tokenRepo")
	}

	return GetStatusSubscriberHandler{
		tokenRepo: tokenRepo,
	}
}

func (c GetStatusSubscriberHandler) Handle(ctx context.Context, query string) (bool, error) {
	hasToken, err := c.tokenRepo.GetToken(query)
	if err != nil {
		return hasToken, errors.NewSlugError(err.Error(), "unable to get token")
	}

	return hasToken, nil
}