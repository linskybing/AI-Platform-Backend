package scheduler

import (
	"context"
	"log"
	"time"

	"github.com/linskybing/platform-go/internal/domain/job"
	"github.com/linskybing/platform-go/internal/scheduler/executor"
	"github.com/linskybing/platform-go/internal/scheduler/queue"
)

// Scheduler manages job execution with priority queue
type Scheduler struct {
	jobQueue *queue.JobQueue
	registry *executor.ExecutorRegistry
	running  bool
}

// NewScheduler creates a new scheduler
func NewScheduler(registry *executor.ExecutorRegistry) *Scheduler {
	return &Scheduler{
		jobQueue: queue.NewJobQueue(),
		registry: registry,
		running:  false,
	}
}

// Start begins scheduling
func (s *Scheduler) Start(ctx context.Context) error {
	s.running = true
	log.Println("Scheduler started")

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			s.running = false
			log.Println("Scheduler stopped")
			return ctx.Err()
		case <-ticker.C:
			s.processQueue(ctx)
		}
	}
}

// EnqueueJob adds a job to queue
func (s *Scheduler) EnqueueJob(j *job.Job) {
	s.jobQueue.Push(j)
}

// processQueue processes pending jobs
func (s *Scheduler) processQueue(ctx context.Context) {
	j := s.jobQueue.Pop()
	if j == nil {
		return
	}

	err := s.registry.Execute(ctx, j)
	if err == executor.ErrExecutorNotFound {
		// Don't change status for unregistered job types
		log.Printf("Job executor not found for type: %s", j.JobType)
		return
	}
	if err != nil {
		log.Printf("Job error: %v", err)
		j.Status = string(job.StatusFailed)
	} else {
		j.Status = string(job.StatusRunning)
	}
}

// IsRunning returns if active
func (s *Scheduler) IsRunning() bool {
	return s.running
}

// GetQueueSize returns pending count
func (s *Scheduler) GetQueueSize() int {
	return s.jobQueue.Len()
}
