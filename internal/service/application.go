package service

import (
	"context"
	"os"

	"github.com/jerensl/jerens-web-api/internal/adapters"
	"github.com/jerensl/jerens-web-api/internal/app"
	"github.com/jerensl/jerens-web-api/internal/app/command"
	"github.com/jerensl/jerens-web-api/internal/app/query"
)

func NewApplication(ctx context.Context) app.Application {
	db, err := adapters.NewSQLiteConnection(os.Getenv("SQLITE_DB"))
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
			Unsubscribe: command.NewUnsubscribe(tokenRepository),
			SendNotification: command.NewSendNotificationHandler(&messaging),
		},
		Queries: app.Queries{
			GetStatusSubscriber: query.NewGetStatusSubscriberHandler(tokenRepository),
			GetAllSubscriber: query.NewGetAllSubscriberHandler(tokenRepository),
		},
	}
}