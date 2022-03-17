package ports

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jerensl/jerens-web-api/internal/app"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(application app.Application) HttpServer {
	return HttpServer{
		app: application,
	}
}

func (h HttpServer) Subscribe(w http.ResponseWriter, r *http.Request) {
	var newSubscriber NewSubscriber
	if err := json.NewDecoder(r.Body).Decode(&newSubscriber); err != nil {
		return
	}

	fmt.Println(newSubscriber)
	w.WriteHeader(http.StatusCreated)
}