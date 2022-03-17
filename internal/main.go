package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jerensl/jerens-web-api/internal/ports"
	"github.com/jerensl/jerens-web-api/internal/server"
	"github.com/jerensl/jerens-web-api/internal/service"
)

func main() {
	ctx := context.Background()

	application := service.NewApplication(ctx)

	server.RunHTTPServer(func(router chi.Router) http.Handler {
			return ports.HandlerFromMux(
				ports.NewHttpServer(application),
				router,
			)
	})
}