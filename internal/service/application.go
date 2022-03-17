package service

import (
	"context"

	"github.com/jerensl/jerens-web-api/internal/app"
)

func NewApplication(ctx context.Context) app.Application {


	return app.Application{
		Commands: app.Commands{
		},
		Queries: app.Queries{
		},
	}
}