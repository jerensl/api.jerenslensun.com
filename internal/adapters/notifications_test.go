package adapters_test

import (
	"testing"

	"github.com/jerensl/api.jerenslensun.com/internal/adapters"
	"github.com/stretchr/testify/assert"
)

func TestFCMConnection(t *testing.T) {
	_, err := adapters.NewFirebaseMessagingConnection()
	assert.NoError(t, err, "cannot connect to FCM")
}

func TestSendMesageOnFCM(t *testing.T) {
	fcm, err := adapters.NewFirebaseMessagingConnection()
	assert.NoError(t, err, "cannot connect to FCM")

	msg := adapters.Messaging{
		MessagingClient: fcm,
	}

	err = msg.SendNotification([]string{"abc"}, "Test Title Message", "Test Body Message")
	assert.NoError(t, err, "Message cannot be send")
}