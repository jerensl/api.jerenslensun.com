package command

import "context"

type NotificationService interface {
	SendNotification(ctx context.Context, token []string, title string, message string) error
}