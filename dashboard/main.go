package main

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"my-asynq/utils"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hibiken/asynqmon"
)

func main() {
	ctx := context.Background()
	h := asynqmon.New(asynqmon.Options{
		RootPath:     "/", // RootPath specifies the root for asynqmon app
		RedisConnOpt: utils.GetRedisClientOpt(ctx),
	})

	r := mux.NewRouter()
	r.PathPrefix(h.RootPath()).Handler(h)

	srv := &http.Server{
		Handler: r,
		Addr:    ":8081",
	}

	// Go to http://localhost:8080/monitoring to see asynqmon homepage.
	g.Log().Panic(ctx, srv.ListenAndServe())
}
