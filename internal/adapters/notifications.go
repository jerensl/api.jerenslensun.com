package adapters

import (
	"context"
	"errors"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

func NewFirebaseMessagingConnection(ctx context.Context) (*messaging.Client, error) {
	var opts []option.ClientOption

	if file := os.Getenv("SERVICE_ACCOUNT_FILE"); file != "" {
		opts = append(opts, option.WithCredentialsFile(os.Getenv("SERVICE_ACCOUNT_FILE")))
	}
	config := &firebase.Config{ProjectID: os.Getenv("PROJECT_ID")}
	firebaseApp, err := firebase.NewApp(context.Background(), config, opts...)
	if err != nil {
		logrus.Fatalf("error initializing app: %v\n", err)
	}

	client, err := firebaseApp.Messaging(ctx)
	if err != nil {
		logrus.Fatalf("error getting Messaging client: %v\n", err)
	}

	return client, nil
}

type Messaging struct {
	MessagingClient *messaging.Client
}

func (m *Messaging) SendNotification(ctx context.Context,token []string, title string, message string) error {
	if len(token) < 1 {
		return errors.New("Unable to get token list")
	}

	notification := &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: title,
			Body: message,
			ImageURL: "https://ik.imagekit.io/jerensl/logo.png",
		},
		Webpush: &messaging.WebpushConfig{
			FCMOptions: &messaging.WebpushFCMOptions{
				Link: "https://www.jerenslensun.com/",
			},
		},
		Tokens: token,
	}

	_, err := m.MessagingClient.SendEachForMulticast(ctx, notification)
	if err != nil {
		return errors.New("Unable to get send notification")
	}

	return nil
}