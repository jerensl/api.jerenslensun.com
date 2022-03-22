package query

import (
	"context"

	"github.com/jerensl/jerens-web-api/internal/logs/errors"
)

type GetStatusSubscriberHandler struct {
	readToModel GetTokenReadModel
}

type GetTokenReadModel interface {
	GetToken(ctx context.Context, token string) (bool, error)
}


func NewGetStatusSubscriberHandler(tokenRepo GetTokenReadModel) GetStatusSubscriberHandler {
	if tokenRepo == nil {
		panic("nil tokenRepo")
	}

	return GetStatusSubscriberHandler{
		readToModel: tokenRepo,
	}
}

func (c GetStatusSubscriberHandler) Handle(ctx context.Context, query string) (bool, error) {
	hasToken, err := c.readToModel.GetToken(ctx, query)
	if err != nil {
		return hasToken, errors.NewSlugError(err.Error(), "unable to get token")
	}

	return hasToken, nil
}