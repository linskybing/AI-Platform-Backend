package priority

import (
	"context"
	"log"
	"sync"

	"github.com/linskybing/platform-go/internal/domain/job"
)

// Manager manages job priority and preemption
type Manager struct {
	mu          sync.RWMutex
	runningJobs map[uint]*job.Job
}

func priorityWeight(p string) int {
	switch p {
	case "high":
		return 3
	case "medium":
		return 2
	default:
		return 1
	}
}

// NewManager creates priority manager
func NewManager() *Manager {
	return &Manager{
		runningJobs: make(map[uint]*job.Job),
	}
}

// RegisterRunningJob registers job
func (m *Manager) RegisterRunningJob(j *job.Job) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.runningJobs[j.ID] = j
}

// UnregisterJob removes job
func (m *Manager) UnregisterJob(jobID uint) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.runningJobs, jobID)
}

// CheckPreemption checks if preemptible
func (m *Manager) CheckPreemption(ctx context.Context, j *job.Job) bool {
	return priorityWeight(j.Priority) < priorityWeight("high")
}

// GetRunningJobs returns running jobs
func (m *Manager) GetRunningJobs() []*job.Job {
	m.mu.RLock()
	defer m.mu.RUnlock()

	jobs := make([]*job.Job, 0, len(m.runningJobs))
	for _, j := range m.runningJobs {
		jobs = append(jobs, j)
	}
	return jobs
}

// CanPreempt checks if preemption allowed
func (m *Manager) CanPreempt(current, incoming *job.Job) bool {
	return priorityWeight(incoming.Priority) > priorityWeight(current.Priority)
}

// PreemptJob preempts job
func (m *Manager) PreemptJob(ctx context.Context, jobID uint) error {
	m.mu.Lock()
	_, exists := m.runningJobs[jobID]
	m.mu.Unlock()

	if !exists {
		return ErrJobNotRunning
	}

	log.Printf("Preempting job %d", jobID)
	m.UnregisterJob(jobID)
	return nil
}
