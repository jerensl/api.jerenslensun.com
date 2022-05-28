package command

import (
	"github.com/jerensl/api.jerenslensun.com/internal/domain"
	"github.com/jerensl/api.jerenslensun.com/internal/logs/errors"
)

type UnsubscribeHandler struct {
	tokenRepo domain.Repository
}

func NewUnsubscribe(tokenRepo domain.Repository) UnsubscribeHandler {
	if tokenRepo == nil {
		panic("nil tokenRepo")
	}

	return UnsubscribeHandler{
		tokenRepo: tokenRepo,
	}
}

func (u UnsubscribeHandler) Handle(token string) error {
	err := u.tokenRepo.DeleteToken(token)
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable to unsubscribe")
	}
	
	return nil
}