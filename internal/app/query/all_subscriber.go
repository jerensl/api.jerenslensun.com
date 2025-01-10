package query

import (
	"context"

	"github.com/jerensl/api.jerenslensun.com/internal/decorator"
	"github.com/jerensl/api.jerenslensun.com/internal/domain/notification"
	"github.com/jerensl/api.jerenslensun.com/internal/logs/errors"
	"github.com/sirupsen/logrus"
)

type AllSubscriber struct {}

type AllSubscriberHandler decorator.QueryHandler[AllSubscriber, []string]

type allSubscriberHandler struct {
	tokenRepo notification.Repository
}

func NewAllSubscriberHandler(tokenRepo notification.Repository, logger *logrus.Entry, metricsClient decorator.MetricsClient) AllSubscriberHandler {
	if tokenRepo == nil {
		panic("nil tokenRepo")
	}

	return decorator.ApplyQueryDecorator[AllSubscriber, []string](
		allSubscriberHandler{tokenRepo: tokenRepo},
		logger,
		metricsClient,
	)
}

func (c allSubscriberHandler) Handle(ctx context.Context, _ AllSubscriber) ([]string, error) {
	subscriber, err := c.tokenRepo.GetAllToken()

	if err != nil {
		return nil, errors.NewSlugError(err.Error(), "unable to get all token")
	}

	return subscriber, nil
}