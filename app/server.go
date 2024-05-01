package main

import (
	"context"
	"log"
	"my-asynq/utils"
	"time"

	"github.com/hibiken/asynq"

	"my-asynq/tasks"
)

func main() {
	ctx := context.Background()

	// Example 1: enqueue task to be processed immediately
	// Use (*Client).Enqueue method
	task, err := tasks.NewEmailDeliveryTask(42, "some:template:id")
	if err != nil {
		log.Fatalf("could not create task %v", err)
	}

	info, err := utils.EnqueueTask(ctx, task)
	if err != nil {
		log.Fatalf("could not enqueue task %v", err)
	}

	log.Printf("enqueued task: id=%s, queue=%s", info.ID, info.Queue)

	// Example 2: enqueue task to be processed in the future
	// Use ProcessIn or ProcessAt option
	info, err = utils.EnqueueTask(ctx, task, asynq.ProcessIn(24*time.Second))
	if err != nil {
		log.Fatalf("could not enqueue task %v", err)
	}
	log.Printf("enqueued task: id=%s, queue=%s", info.ID, info.Queue)

	// Example 3: set other options to tune task processing behavior
	// Options include MaxRetry, Queue, Timeout, Deadline, Unique etc
	task, err = tasks.NewImageResizeTask("https://example.com/myassets/image.jpg")
	info, err = utils.EnqueueTask(ctx, task, asynq.MaxRetry(10), asynq.Timeout(3*time.Minute))
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
}
