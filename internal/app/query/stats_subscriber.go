package query

import (
	"context"

	"github.com/jerensl/api.jerenslensun.com/internal/decorator"
	"github.com/jerensl/api.jerenslensun.com/internal/domain/notification"
	"github.com/jerensl/api.jerenslensun.com/internal/logs/errors"
	"github.com/sirupsen/logrus"
)

type StatsSubscriber struct {}

type StatsSubscriberHandler decorator.QueryHandler[StatsSubscriber, *notification.Stats]

type statsSubscriberHandler struct {
	tokenRepo notification.Repository
}

func NewStatsSubscriberHandler(tokenRepo notification.Repository, logger *logrus.Entry, metricsClient decorator.MetricsClient) StatsSubscriberHandler {
	if tokenRepo == nil {
		panic("nil tokenRepo")
	}

	return decorator.ApplyQueryDecorator[StatsSubscriber, *notification.Stats](
		statsSubscriberHandler{
			tokenRepo: tokenRepo,
		},
		logger,
		metricsClient,
	)
}

func (c statsSubscriberHandler) Handle(_ context.Context, query StatsSubscriber) (*notification.Stats, error) {
	stats, err := c.tokenRepo.GetStatisticToken()

	if err != nil {
		return nil, errors.NewSlugError(err.Error(), "unable to get stats token")
	}

	return stats, nil
}