package notification_test

import (
	"testing"
	"time"

	"github.com/jerensl/api.jerenslensun.com/internal/domain/notification"
	"github.com/stretchr/testify/require"
)

func TestNewToken(t *testing.T) {
	tokenID := "ab2941j4149j"
	isActive := true
	updatedAt := time.Now().Unix()
	token, err := notification.NewToken(tokenID, isActive, time.Now().Unix())
	require.NoError(t, err)

	require.Equal(t, tokenID, token.TokenID())
	require.Equal(t, isActive, token.IsActive())
	require.Equal(t, time.Unix(updatedAt, 0), token.UpdatedAt())
}

func TestNewTokenInvalid(t *testing.T) {
	tokenID := "ab2941j4149j"
	isActive := true
	updatedAt := time.Now().Unix()
	_, err := notification.NewToken("", isActive, updatedAt)
	require.Error(t, err)

	_, err = notification.NewToken(tokenID, isActive, 0)
	require.Error(t, err)
}