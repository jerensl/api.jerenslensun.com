package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jerensl/api.jerenslensun.com/internal/ports"
	"github.com/jerensl/api.jerenslensun.com/internal/server"
	"github.com/jerensl/api.jerenslensun.com/internal/service"
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