package priority

import (
	"context"
	"testing"

	"github.com/linskybing/platform-go/internal/domain/job"
)

func TestRegisterAndPreemptJob(t *testing.T) {
	m := NewManager()
	j := &job.Job{ID: 1, Priority: "low"}
	m.RegisterRunningJob(j)

	if len(m.GetRunningJobs()) != 1 {
		t.Fatalf("expected 1 running job")
	}

	if err := m.PreemptJob(context.Background(), j.ID); err != nil {
		t.Fatalf("unexpected error preempting job: %v", err)
	}

	if len(m.GetRunningJobs()) != 0 {
		t.Fatalf("expected no running jobs after preemption")
	}

	if err := m.PreemptJob(context.Background(), 99); err != ErrJobNotRunning {
		t.Fatalf("expected ErrJobNotRunning, got %v", err)
	}
}

func TestCanPreempt(t *testing.T) {
	m := NewManager()
	high := &job.Job{ID: 1, Priority: "high"}
	low := &job.Job{ID: 2, Priority: "low"}

	if m.CanPreempt(high, low) {
		t.Fatalf("high priority job should not be preempted by low priority job")
	}

	if !m.CanPreempt(low, high) {
		t.Fatalf("low priority job should be preempted by high priority job")
	}
}

func TestCheckPreemption(t *testing.T) {
	m := NewManager()
	high := &job.Job{Priority: "high"}
	low := &job.Job{Priority: "low"}

	if m.CheckPreemption(context.Background(), high) {
		t.Fatalf("high priority job should not be preemptible")
	}

	if !m.CheckPreemption(context.Background(), low) {
		t.Fatalf("low priority job should be preemptible")
	}
}
