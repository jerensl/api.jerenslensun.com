package query

import (
	"context"

	"github.com/jerensl/jerens-web-api/internal/logs/errors"
)

type GetAllSubscriberHandler struct {
	readToModel GetAllSubscriberReadModel
}

type GetAllSubscriberReadModel interface {
	GetAllToken(ctx context.Context) ([]string, error)
}


func NewGetAllSubscriberHandler(AllSubscriberRepo GetAllSubscriberReadModel) GetAllSubscriberHandler {
	if AllSubscriberRepo == nil {
		panic("nil AllSubscriberRepo")
	}

	return GetAllSubscriberHandler{
		readToModel: AllSubscriberRepo,
	}
}

func (c GetAllSubscriberHandler) Handle(ctx context.Context) ([]string, error) {
	subscriber, err := c.readToModel.GetAllToken(ctx)

	if err != nil {
		return nil, errors.NewSlugError(err.Error(), "unable to get all token")
	}

	return subscriber, nil
}