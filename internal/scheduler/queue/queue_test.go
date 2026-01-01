package queue

import (
	"testing"

	"github.com/linskybing/platform-go/internal/domain/job"
)

func TestJobQueuePriorityOrder(t *testing.T) {
	q := NewJobQueue()

	low := &job.Job{ID: 1, Priority: "low"}
	medium := &job.Job{ID: 2, Priority: "medium"}
	high := &job.Job{ID: 3, Priority: "high"}

	q.Push(low)
	q.Push(medium)
	q.Push(high)

	if q.Len() != 3 {
		t.Fatalf("expected queue length 3, got %d", q.Len())
	}

	if got := q.Pop(); got.ID != high.ID {
		t.Fatalf("expected high priority first, got %d", got.ID)
	}
	if got := q.Pop(); got.ID != medium.ID {
		t.Fatalf("expected medium priority second, got %d", got.ID)
	}
	if got := q.Pop(); got.ID != low.ID {
		t.Fatalf("expected low priority last, got %d", got.ID)
	}

	if !q.IsEmpty() {
		t.Fatalf("expected queue to be empty")
	}
}

func TestJobQueuePeek(t *testing.T) {
	q := NewJobQueue()
	high := &job.Job{ID: 9, Priority: "high"}

	q.Push(high)
	peek := q.Peek()
	if peek == nil || peek.ID != high.ID {
		t.Fatalf("expected peek to return job %d", high.ID)
	}

	if q.Len() != 1 {
		t.Fatalf("expected length to remain 1 after peek, got %d", q.Len())
	}
}
