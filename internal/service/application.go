package service

import (
	"context"

	"github.com/jerensl/jerens-web-api/internal/adapters"
	"github.com/jerensl/jerens-web-api/internal/app"
	"github.com/jerensl/jerens-web-api/internal/app/command"
	"github.com/jerensl/jerens-web-api/internal/app/query"
)

func NewApplication(ctx context.Context) app.Application {
	db, err := adapters.NewSQLiteConnection("../database/sqlite.db")
	if err != nil {
		panic(err)
	}

	tokenRepository := adapters.NewSQLiteTokenRepository(db)
	messageClient, err := adapters.NewFirebaseMessagingConnection()
	if err != nil {
		panic(err)
	}

	messaging := adapters.Messaging{
		MessagingClient: messageClient,
	}

	return app.Application{
		Commands: app.Commands{
			AddNewSubscriber: command.NewAddNewSubscriberHandler(tokenRepository),
			SendNotification: command.NewSendNotificationHandler(&messaging),
		},
		Queries: app.Queries{
			CheckIfTokenExist: query.NewCheckTokenHandler(tokenRepository),
			GetAllSubscriber: query.NewGetAllSubscriberHandler(tokenRepository),
		},
	}
}