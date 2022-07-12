package command

import (
	"github.com/jerensl/api.jerenslensun.com/internal/domain/notification"
	"github.com/jerensl/api.jerenslensun.com/internal/logs/errors"
)

type UnsubscribeHandler struct {
	tokenRepo notification.Repository
}

func NewUnsubscribe(tokenRepo notification.Repository) UnsubscribeHandler {
	if tokenRepo == nil {
		panic("nil tokenRepo")
	}

	return UnsubscribeHandler{
		tokenRepo: tokenRepo,
	}
}

func (u UnsubscribeHandler) Handle(tokenID string, updateAt int64) (*notification.Token, error) {
	token, err := u.tokenRepo.UpdatedToken(tokenID, updateAt)
	if err != nil {
		return token, errors.NewSlugError(err.Error(), "unable to unsubscribe")
	}
	
	return token, nil
}