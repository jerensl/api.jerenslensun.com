package service

import (
	"context"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/jerensl/jerens-web-api/internal/ports"
	"github.com/jerensl/jerens-web-api/internal/server"
	"github.com/jerensl/jerens-web-api/internal/tests"
)


func TestGetAllSubscriber(t *testing.T) {
	client := tests.NewAppHttpClient(t)
	token := "abc"

	client.SubscriberStatus(t, token)
}


func startService() bool {
	app := NewApplication(context.Background())

	appHTTPAddr := os.Getenv("APP_HTTP_ADDR")

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

	os.Exit(m.Run())
}