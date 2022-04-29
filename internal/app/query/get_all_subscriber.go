package query

import (
	"context"

	"github.com/jerensl/api.jerenslensun.com/internal/logs/errors"
)

type GetAllSubscriberHandler struct {
	readToModel GetAllSubscriberReadModel
}

type GetAllSubscriberReadModel interface {
	GetAllToken() ([]string, error)
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
	subscriber, err := c.readToModel.GetAllToken()

	if err != nil {
		return nil, errors.NewSlugError(err.Error(), "unable to get all token")
	}

	return subscriber, nil
}