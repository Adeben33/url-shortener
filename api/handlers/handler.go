package handlers

import (
	"context"
	"fmt"
	"net/http"
)

type Handler struct {
	ctx context.Context
}

type handler interface {
	CreateUrl(w http.ResponseWriter, r *http.Request)
}

func APiHandler(ctx context.Context) handler {
	return &Handler{ctx: ctx}
}

func (i *Handler) CreateUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}
