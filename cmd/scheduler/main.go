package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/linskybing/platform-go/internal/application/scheduler"
	"github.com/linskybing/platform-go/internal/scheduler/executor"
)

func main() {
	registry := executor.NewExecutorRegistry()
	sched := scheduler.NewScheduler(registry)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
		<-sigChan
		log.Println("Shutdown signal")
		cancel()
	}()

	log.Printf("Starting scheduler (queue: %d)", sched.GetQueueSize())
	if err := sched.Start(ctx); err != nil {
		log.Printf("Scheduler error: %v", err)
	}
}
