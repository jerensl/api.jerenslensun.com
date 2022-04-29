package command

import "github.com/jerensl/api.jerenslensun.com/internal/logs/errors"

type UnsubscribeHandler struct {
	writeToModel UnsubscriberWriteModel
}

type UnsubscriberWriteModel interface {
	DeleteToken(token string) error
}

func NewUnsubscribe(tokenRepo UnsubscriberWriteModel) UnsubscribeHandler {
	if tokenRepo == nil {
		panic("nil tokenRepo")
	}

	return UnsubscribeHandler{
		writeToModel: tokenRepo,
	}
}

func (u UnsubscribeHandler) Handle(token string) error {
	err := u.writeToModel.DeleteToken(token)
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable to unsubscribe")
	}
	
	return nil
}