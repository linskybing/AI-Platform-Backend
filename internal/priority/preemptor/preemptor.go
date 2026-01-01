package preemptor

import (
	"context"
	"time"
)

// PreemptionDecision represents a decision to preempt jobs
type PreemptionDecision struct {
	JobsToPreempt    []uint
	Reason           string
	RequestedBy      uint
	EstimatedTime    time.Time
	GracePeriod      time.Duration
	CreateCheckpoint bool
}

// Preemptor defines the interface for job preemption
type Preemptor interface {
	SelectJobsToPreempt(ctx context.Context, requiredResources ResourceRequirement) (*PreemptionDecision, error)
	PreemptJob(ctx context.Context, jobID uint) error
	CreateCheckpoint(ctx context.Context, jobID uint) error
	NotifyPreemption(ctx context.Context, jobID uint, reason string) error
}

// ResourceRequirement represents resource requirements for a workload
type ResourceRequirement struct {
	CPU    float64
	Memory int64
	GPU    int
	MPS    int
}

// PreemptionConfig defines preemption behavior configuration
type PreemptionConfig struct {
	GracePeriod      time.Duration
	EnableCheckpoint bool
	MaxPreemptions   int
	NotifyUsers      bool
}

// DefaultPreemptionConfig returns default preemption configuration
func DefaultPreemptionConfig() *PreemptionConfig {
	return &PreemptionConfig{
		GracePeriod:      120 * time.Second,
		EnableCheckpoint: true,
		MaxPreemptions:   5,
		NotifyUsers:      true,
	}
}
