package service

import (
	"context"

	"github.com/jerensl/api.jerenslensun.com/internal/adapters"
	"github.com/jerensl/api.jerenslensun.com/internal/app"
	"github.com/jerensl/api.jerenslensun.com/internal/app/command"
	"github.com/jerensl/api.jerenslensun.com/internal/app/query"
)

func NewApplication(ctx context.Context) app.Application {
	db, err := adapters.NewSQLiteConnection()
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

	return app.Application{
		Commands: app.Commands{
			AddNewSubscriber: command.NewAddNewSubscriberHandler(tokenRepository),
			Unsubscribe: command.NewUnsubscribe(tokenRepository),
			SendNotification: command.NewSendNotificationHandler(&messaging),
		},
		Queries: app.Queries{
			GetStatusSubscriber: query.NewGetStatusSubscriberHandler(tokenRepository),
			GetAllSubscriber: query.NewGetAllSubscriberHandler(tokenRepository),
		},
	}
}