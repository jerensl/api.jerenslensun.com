package service

import (
	"context"
	"os"

	"github.com/jerensl/api.jerenslensun.com/internal/adapters"
	"github.com/jerensl/api.jerenslensun.com/internal/app"
	"github.com/jerensl/api.jerenslensun.com/internal/app/command"
	"github.com/jerensl/api.jerenslensun.com/internal/app/query"
	"github.com/jerensl/api.jerenslensun.com/internal/metrics"
	"github.com/sirupsen/logrus"
)

func NewApplication(ctx context.Context) app.Application {
	db, err := adapters.NewSQLiteConnection(os.Getenv("SQLITE_DB"))
	if err != nil {
		panic(err)
	}

	tokenRepository := adapters.NewSQLiteTokenRepository(db)
	messageClient, err := adapters.NewFirebaseMessagingConnection(ctx)
	if err != nil {
		panic(err)
	}

	messaging := adapters.Messaging{
		MessagingClient: messageClient,
	}

	logger := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := metrics.NoOp{}

	return app.Application{
		Commands: app.Commands{
			AddNewSubscriber: command.NewAddNewSubscriberHandler(tokenRepository, logger, metricsClient),
			Unsubscribe: command.NewUnsubscribe(tokenRepository, logger, metricsClient),
			SendNotification: command.NewSendNotificationHandler(&messaging),
		},
		Queries: app.Queries{
			GetStatusSubscriber: query.NewGetStatusSubscriberHandler(tokenRepository),
			GetAllSubscriber: query.NewGetAllSubscriberHandler(tokenRepository),
			GetStatsSubscriber: query.NewGetStatsSubscriberHandler(tokenRepository),
		},
	}
}