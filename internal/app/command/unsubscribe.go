package command

import (
	"context"

	"github.com/jerensl/api.jerenslensun.com/internal/decorator"
	"github.com/jerensl/api.jerenslensun.com/internal/domain/notification"
	"github.com/jerensl/api.jerenslensun.com/internal/logs/errors"
	"github.com/sirupsen/logrus"
)

type Unsubscribe struct {
	TokenID		string
	UpdateAt	int64
}

type UnsubscribeHandler decorator.CommandHandler[Unsubscribe]

type unsubscribeHandler struct {
	tokenRepo notification.Repository
}

func NewUnsubscribe(tokenRepo notification.Repository, logger *logrus.Entry, metricsClient decorator.MetricsClient) UnsubscribeHandler {
	if tokenRepo == nil {
		panic("nil tokenRepo")
	}

	return decorator.ApplyCommandDecorator[Unsubscribe](
		unsubscribeHandler{tokenRepo: tokenRepo},
		logger,
		metricsClient,
	)
}

func (u unsubscribeHandler) Handle(ctx context.Context, cmd Unsubscribe) (error) {
	_, err := u.tokenRepo.UpdatedToken(cmd.TokenID, cmd.UpdateAt)
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable to unsubscribe")
	}
	
	return nil
}