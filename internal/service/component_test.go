package service

import (
	"context"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/jerensl/api.jerenslensun.com/internal/ports"
	"github.com/jerensl/api.jerenslensun.com/internal/server"
	"github.com/jerensl/api.jerenslensun.com/internal/tests"
)



func TestGetStatusNotSubscriber(t *testing.T) {
	client := tests.NewHttpClient(t)
	token := "abc"

	client.NotSubscriberStatus(t, token)
}

func TestSubscribeNotification(t *testing.T) {
	client := tests.NewHttpClient(t)
	token := "abc"

	client.SubscibeNotification(t, token)
}

func TestAlreadySubscribeNotification(t *testing.T) {
	client := tests.NewHttpClient(t)
	token := "abc"

	client.AlreadySubscibeNotification(t, token)
}

func TestGetStatusAlreadySubscriber(t *testing.T) {
	client := tests.NewHttpClient(t)
	token := "abc"

	client.AlreadySubscriberStatus(t, token)
}

func TestSendNotification(t *testing.T) {
	client := tests.NewHttpClient(t)
	title := "Test title"
	message := "Test Message"

	client.SendNotification(t, title, message)
}

func TestSendNotificationWithouAutz(t *testing.T) {
	client := tests.NewHttpClient(t)
	title := "Test title"
	message := "Test Message"

	client.SendNotificationWithoutAuthz(t, title, message)
}

func TestUnsubscribeNotification(t *testing.T) {
	client := tests.NewHttpClient(t)
	token := "abc"

	client.UnsubscibeNotification(t, token)
}

func TestUnsubscribeFromSubsriberNotExist(t *testing.T) {
	client := tests.NewHttpClient(t)
	token := "abc123"

	client.UnsubsciberFromSubscriberNotExist(t, token)
}


func startService() bool {
	app := NewApplication(context.Background())

	appHTTPAddr := os.Getenv("HTTP_ADDR")

	go server.RunHTTPServerOnAddr(appHTTPAddr, func(router chi.Router) http.Handler {
		return ports.HandlerFromMux(ports.NewHttpServer(app), router)
	})

	ok := tests.WaitForPort(appHTTPAddr)
	if !ok {
		log.Println("Timed out waiting for app HTTP to come out")
		return false
	}

	return ok
}

func TestMain(m *testing.M) {
	if !startService() {
		log.Println("Timed out waiting for HTTP to come out")
		os.Exit(1)
	}

	testCases := m.Run()
	os.Remove(os.Getenv("SQLITE_DB"))
	os.Exit(testCases)
}