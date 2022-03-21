package query

import (
	"context"

	"github.com/jerensl/jerens-web-api/internal/logs/errors"
)

type CheckTokenHandler struct {
	readToModel CheckTokenReadModel
}

type CheckTokenReadModel interface {
	GetToken(ctx context.Context, token string) (bool, error)
}


func NewCheckTokenHandler(tokenRepo CheckTokenReadModel) CheckTokenHandler {
	if tokenRepo == nil {
		panic("nil tokenRepo")
	}

	return CheckTokenHandler{
		readToModel: tokenRepo,
	}
}

func (c CheckTokenHandler) Handle(ctx context.Context, query string) (bool, error) {
	hasToken, err := c.readToModel.GetToken(ctx, query)
	if err != nil {
		return hasToken, errors.NewSlugError(err.Error(), "unable to get token")
	}

	return hasToken, nil
}