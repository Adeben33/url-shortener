package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
	"log"
	"math/rand"
	"net/http"
	"time"
	"url-shortener/models"
	redisRepo "url-shortener/repo/redis"
	"url-shortener/utils"
)

type Handler struct {
	ctx     context.Context
	redisDb redisRepo.RedisInterface
}

type handler interface {
	ShortenUrl(w http.ResponseWriter, r *http.Request)
	UrlRedirect(w http.ResponseWriter, r *http.Request)
	PingTheDb(w http.ResponseWriter, r *http.Request)
	Ping(w http.ResponseWriter, r *http.Request)
}

func APiHandler(ctx context.Context, redisInterface redisRepo.RedisInterface) handler {
	return &Handler{ctx: ctx,
		redisDb: redisInterface}
}

func (i *Handler) ShortenUrl(w http.ResponseWriter, r *http.Request) {
	var req models.UrlReq
	var identifier string
	var exp time.Duration

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	log.Println(req.CustomShort)

	if req.CustomShort == " " {
		identifier = utils.Base62Converter(rand.Uint64())
	} else if req.CustomShort != " " {
		identifier = req.CustomShort
	}
	log.Println(identifier)
	if req.Expiry == 0 {
		exp = 24
	}
	_, err = i.redisDb.Get(identifier)
	if err == nil {
		msg := "URL Custom short is already in use"
		w.Write([]byte(msg))
		w.WriteHeader(http.StatusNotFound)
		return
	}
	//save to redis
	result, err := i.redisDb.Set(identifier, req.URL, exp*3600*time.Second)
	if err == redis.Nil {
		msg := "URL Custom short is already in use"
		w.Write([]byte(msg))
		w.WriteHeader(404)
		return
	}
	msg := map[string]string{
		"msg":    "custom url creation successful",
		"url":    fmt.Sprintf("http://localhost:8080/%v", identifier),
		"result": *result,
	}
	msgjson, _ := json.Marshal(msg)
	w.Write(msgjson)
	w.WriteHeader(200)
}

func (i *Handler) PingTheDb(w http.ResponseWriter, r *http.Request) {
	result1, _ := i.redisDb.Ping()
	fmt.Fprint(w, *result1)
}

func (i *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ping")
}

func (i *Handler) UrlRedirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := vars["url"]

	result, err := i.redisDb.Get(url)
	if err == redis.Nil {
		msg := "url does not exist"
		w.WriteHeader(404)
		fmt.Fprint(w, msg)
		return
	} else if err != nil {
		msg := "error found"
		w.WriteHeader(404)
		fmt.Fprintf(w, msg, err.Error())
		return
	}

	http.Redirect(w, r, *result, http.StatusSeeOther)
}

func (i *Handler) UrlDetail(w http.ResponseWriter, r *http.Request) {

}
