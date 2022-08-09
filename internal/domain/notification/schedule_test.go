package notification_test

import (
	"testing"
	"time"

	"github.com/jerensl/api.jerenslensun.com/internal/domain/notification"
	"github.com/stretchr/testify/assert"
)


func TestNewScheduler(t *testing.T) {
	result := make(map[string]bool)
	scheduler := notification.NewScheduler(5, func(title, message string) {
		result[title] = true
	})

	newTitle := "hello"
	newMsg := "new messages"

	scheduler.NewJob("Test", newTitle, newMsg, time.Second * 5)
	time.Sleep(time.Second * 5)
	assert.True(t, result[newTitle])
}

func TestNewSchedulerLimit(t *testing.T) {
	result := make(map[string]bool)
	scheduler := notification.NewScheduler(2, func(title, message string) {
		result[title] = true
	})

	newTitle := "hello"
	newMsg := "new messages"

	scheduler.NewJob("Test", newTitle, newMsg, time.Second * 5)
	scheduler.NewJob("Test2", newTitle, newMsg, time.Second * 3)
	time.Sleep(time.Second * 5)
	assert.True(t, result[newTitle])
}