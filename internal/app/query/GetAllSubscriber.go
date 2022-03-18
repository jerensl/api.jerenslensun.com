package query

import (
	"context"

	"github.com/sirupsen/logrus"
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

func (c GetAllSubscriberHandler) Handle(ctx context.Context) (subscriber []string, err error) {
	defer func() {
		logrus.WithError(err).Debug("GetAllSubscriberHandler Executed")
	}()

	return c.readToModel.GetAllToken(ctx)
}