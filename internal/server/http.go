package server

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/sirupsen/logrus"

	"github.com/jerensl/api.jerenslensun.com/internal/logs"
)

func RunHTTPServer(createHandler func(router chi.Router) http.Handler) {
	RunHTTPServerOnAddr(":"+os.Getenv("PORT"), createHandler)
}

func RunHTTPServerOnAddr(addr string, createHandler func(router chi.Router) http.Handler) {
	rootRouter := chi.NewRouter()
	
	apiRouter := chi.NewRouter()
	setMiddlewares(apiRouter)

	rootRouter.Mount("/api", createHandler(apiRouter))

	rootRouter.Mount("/docs", http.StripPrefix("/docs", http.FileServer(http.Dir("./docs"))))
	
	logrus.Info("Starting RESTFull Api server on Port "+os.Getenv("PORT"))

	_ = http.ListenAndServe(addr, rootRouter)
}

func setMiddlewares(router *chi.Mux)  {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(logs.NewStructuredLogger(logrus.StandardLogger()))
	router.Use(httprate.LimitByIP(60, 1*time.Minute))
	router.Use(middleware.Recoverer)

	addCorsMiddleware(router)

	router.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
	)
	router.Use(middleware.NoCache)
}

func addCorsMiddleware(router *chi.Mux) {
	allowedOrigins := strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ";")
	if len(allowedOrigins) == 0 {
		return
	}

	corsMiddleware := cors.Handler(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	})
	
	router.Use(corsMiddleware)
}