package query

import (
	"context"

	"github.com/jerensl/api.jerenslensun.com/internal/decorator"
	"github.com/jerensl/api.jerenslensun.com/internal/domain/notification"
	"github.com/jerensl/api.jerenslensun.com/internal/logs/errors"
	"github.com/sirupsen/logrus"
)

type StatusSubscriber struct {
	TokenID string
}

type StatusSubscriberHandler decorator.QueryHandler[StatusSubscriber, *notification.Token]

type statusSubscriberHandler struct {
	tokenRepo notification.Repository
}

func NewStatusSubscriberHandler(tokenRepo notification.Repository, logger *logrus.Entry, metricsClient decorator.MetricsClient) StatusSubscriberHandler {
	if tokenRepo == nil {
		panic("nil tokenRepo")
	}

	return decorator.ApplyQueryDecorator[StatusSubscriber, *notification.Token](
		statusSubscriberHandler{
			tokenRepo: tokenRepo,
		},
		logger,
		metricsClient,
	)
}

func (c statusSubscriberHandler) Handle(ctx context.Context, query StatusSubscriber) (*notification.Token, error) {
	token, _, err := c.tokenRepo.GetToken(query.TokenID)
	if err != nil {
		return token, errors.NewSlugError(err.Error(), "unable to get token")
	}

	return token, nil
}