package app

import (
	"github.com/jerensl/api.jerenslensun.com/internal/app/command"
	"github.com/jerensl/api.jerenslensun.com/internal/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	AddNewSubscriber command.AddNewSubscriberHandler
	SendNotification command.SendNotificationHandler
	Unsubscribe command.UnsubscribeHandler
}

type Queries struct {
	GetStatusSubscriber query.GetStatusSubscriberHandler
	GetAllSubscriber query.GetAllSubscriberHandler
}