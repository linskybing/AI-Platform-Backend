package executor

import (
	"context"
	"errors"

	"github.com/linskybing/platform-go/internal/domain/job"
)

// ErrExecutorNotFound is returned when an executor is not registered for a job type
var ErrExecutorNotFound = errors.New("executor not found")

// JobExecutor defines the interface for executing different job types
type JobExecutor interface {
	Execute(ctx context.Context, j *job.Job) error
	Cancel(ctx context.Context, jobID uint) error
	GetStatus(ctx context.Context, jobID uint) (job.JobStatus, error)
	GetLogs(ctx context.Context, jobID uint) (string, error)
	SupportsType(jobType job.JobType) bool
}

// ExecutorRegistry manages different job executors
type ExecutorRegistry struct {
	executors map[job.JobType]JobExecutor
}

// NewExecutorRegistry creates a new executor registry
func NewExecutorRegistry() *ExecutorRegistry {
	return &ExecutorRegistry{
		executors: make(map[job.JobType]JobExecutor),
	}
}

// Register registers an executor for a job type
func (r *ExecutorRegistry) Register(jobType job.JobType, executor JobExecutor) {
	r.executors[jobType] = executor
}

// GetExecutor returns the executor for a job type
func (r *ExecutorRegistry) GetExecutor(jobType job.JobType) (JobExecutor, bool) {
	executor, exists := r.executors[jobType]
	return executor, exists
}

// Execute executes a job using the appropriate executor
func (r *ExecutorRegistry) Execute(ctx context.Context, j *job.Job) error {
	executor, exists := r.GetExecutor(j.JobType)
	if !exists {
		return ErrExecutorNotFound
	}
	return executor.Execute(ctx, j)
}
