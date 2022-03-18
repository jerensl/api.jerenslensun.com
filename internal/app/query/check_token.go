package query

import (
	"context"

	"github.com/sirupsen/logrus"
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

func (c CheckTokenHandler) Handle(ctx context.Context, query string) (hasToken bool, err error) {
	defer func() {
		logrus.WithError(err).Debug("CheckTokenHandler Executed")
	}()

	return c.readToModel.GetToken(ctx, query)
}