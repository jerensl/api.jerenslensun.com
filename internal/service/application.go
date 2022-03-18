package service

import (
	"context"

	"github.com/jerensl/jerens-web-api/internal/adapters"
	"github.com/jerensl/jerens-web-api/internal/app"
	"github.com/jerensl/jerens-web-api/internal/app/command"
	"github.com/jerensl/jerens-web-api/internal/app/query"
)

func NewApplication(ctx context.Context) app.Application {
	db, err := adapters.NewSQLiteConnection()
	if err != nil {
		panic(err)
	}


	tokenRepository := adapters.NewSQLiteTokenRepository(db)

	return app.Application{
		Commands: app.Commands{
			AddNewSubscriber: command.NewAddNewSubscriberHandler(tokenRepository),
		},
		Queries: app.Queries{
			CheckIfTokenExist: query.NewCheckTokenHandler(tokenRepository),
		},
	}
}