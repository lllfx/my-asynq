package main

import (
	"context"
	"log"
	"my-asynq/utils"

	"github.com/hibiken/asynq"

	"my-asynq/tasks"
)

func main() {
	ctx := context.Background()
	srv := asynq.NewServer(
		utils.GetRedisClientOpt(ctx),
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			// See the godoc for other configuration options
		},
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	mux.Handle(tasks.TypeEmailDelivery, utils.NewTaskHandler(tasks.HandleEmailDeliveryTask))
	mux.Handle(tasks.TypeImageResize, utils.NewTaskHandler(tasks.ImageProcessorProcessTask))
	// ...register other handlers...

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
