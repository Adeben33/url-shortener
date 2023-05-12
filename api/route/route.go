package route

import (
	"context"
	"github.com/gorilla/mux"
	"url-shortener/api/handlers"
)

func Route(r *mux.Router) *mux.Router {
	ctx := context.Background()
	handler := handlers.APiHandler(ctx)

	r.HandleFunc("/create", handler.CreateUrl).Methods("GET")

	return r
}
