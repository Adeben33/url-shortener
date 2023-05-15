package route

import (
	"context"
	"github.com/gorilla/mux"
	"url-shortener/api/handlers"
	db "url-shortener/database/redis"
	"url-shortener/repo/redis"
)

func Route(r *mux.Router) *mux.Router {
	ctx := context.Background()
	client := db.GetRedisDb()
	redisInterface := redis.RedisRepo(client, ctx)
	handler := handlers.APiHandler(ctx, redisInterface)
	r.HandleFunc("/ping", handler.Ping).Methods("GET")
	r.HandleFunc("/dbping", handler.PingTheDb).Methods("GET")

	r.HandleFunc("/create", handler.ShortenUrl).Methods("POST")
	r.HandleFunc("/{url}", handler.UrlRedirect).Methods("GET")
	r.HandleFunc("/urldetails", handler.UrlRedirect).Methods("GET")

	return r
}
