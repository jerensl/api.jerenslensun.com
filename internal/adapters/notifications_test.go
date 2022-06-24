package adapters_test

import (
	"context"
	"testing"

	"github.com/jerensl/api.jerenslensun.com/internal/adapters"
	"github.com/stretchr/testify/assert"
)

func TestFCMConnection(t *testing.T) {
	ctx := context.Background()
	_, err := adapters.NewFirebaseMessagingConnection(ctx)
	assert.NoError(t, err, "cannot connect to FCM")
}

func TestSendMesageOnFCM(t *testing.T) {
	ctx := context.Background()
	fcm, err := adapters.NewFirebaseMessagingConnection(ctx)
	assert.NoError(t, err, "cannot connect to FCM")

	msg := adapters.Messaging{
		MessagingClient: fcm,
	}

	err = msg.SendNotification(ctx, []string{"abc"}, "Test Title Message", "Test Body Message")
	assert.NoError(t, err, "Message cannot be send")
}

func TestSendMesageOnFCMWithZeroToken(t *testing.T) {
	ctx := context.Background()
	fcm, err := adapters.NewFirebaseMessagingConnection(ctx)
	assert.NoError(t, err, "cannot connect to FCM")

	msg := adapters.Messaging{
		MessagingClient: fcm,
	}

	err = msg.SendNotification(ctx, []string{}, "Test Title Message", "Test Body Message")
	assert.ErrorContains(t, err, "Unable to get token list")
}