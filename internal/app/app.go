package app

import (
	"github.com/jerensl/jerens-web-api/internal/app/command"
	"github.com/jerensl/jerens-web-api/internal/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	AddNewSubscriber command.AddNewSubscriberHandler
}

type Queries struct {
	CheckIfTokenExist query.CheckTokenHandler
}