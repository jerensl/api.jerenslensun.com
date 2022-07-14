package query

import (
	"github.com/jerensl/api.jerenslensun.com/internal/domain/notification"
	"github.com/jerensl/api.jerenslensun.com/internal/logs/errors"
)

type GetStatsSubscriberHandler struct {
	tokenRepo notification.Repository
}


func NewGetStatsSubscriberHandler(tokenRepo notification.Repository) GetStatsSubscriberHandler {
	if tokenRepo == nil {
		panic("nil tokenRepo")
	}

	return GetStatsSubscriberHandler{
		tokenRepo: tokenRepo,
	}
}

func (c GetStatsSubscriberHandler) Handle() (*notification.Stats, error) {
	stats, err := c.tokenRepo.GetStatisticToken()

	if err != nil {
		return nil, errors.NewSlugError(err.Error(), "unable to get stats token")
	}

	return stats, nil
}