package command

import (
	"context"

	"github.com/jerensl/api.jerenslensun.com/internal/decorator"
	"github.com/jerensl/api.jerenslensun.com/internal/domain/notification"
	"github.com/jerensl/api.jerenslensun.com/internal/logs/errors"
	"github.com/sirupsen/logrus"
)

type AddNewSubscriber struct {
	TokenID		string 
	UpdateAt	int64
}

type AddNewSubscriberHandler decorator.CommandHandler[AddNewSubscriber]

type addNewSubscriberHandler struct {
	tokenRepo notification.Repository
}

func NewAddNewSubscriberHandler(tokenRepo notification.Repository, logger *logrus.Entry, metricsClient decorator.MetricsClient) AddNewSubscriberHandler {
	if tokenRepo == nil {
		panic("nil tokenRepo")
	}

	return decorator.ApplyCommandDecorator[AddNewSubscriber](
		addNewSubscriberHandler{tokenRepo: tokenRepo},
		logger,
		metricsClient,
	)
}

func (c addNewSubscriberHandler) Handle(ctx context.Context, cmd AddNewSubscriber) (error) {
	_, err := c.tokenRepo.UpdatedToken(cmd.TokenID, cmd.UpdateAt)
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable to add subscriber")
	}

	return nil
}