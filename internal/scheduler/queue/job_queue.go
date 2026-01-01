package queue

import (
	"container/heap"
	"sync"

	"github.com/linskybing/platform-go/internal/domain/job"
)

// JobQueue is a priority queue for jobs
type JobQueue struct {
	mu    sync.RWMutex
	items priorityQueue
}

// NewJobQueue creates a new job queue
func NewJobQueue() *JobQueue {
	jq := &JobQueue{
		items: make(priorityQueue, 0),
	}
	heap.Init(&jq.items)
	return jq
}

// Push adds a job to the queue
func (jq *JobQueue) Push(j *job.Job) {
	jq.mu.Lock()
	defer jq.mu.Unlock()
	p := 0
	if j.Priority == "high" {
		p = 3
	} else if j.Priority == "medium" {
		p = 2
	} else {
		p = 1
	}
	heap.Push(&jq.items, &queueItem{job: j, priority: p})
}

// Pop removes and returns the highest priority job
func (jq *JobQueue) Pop() *job.Job {
	jq.mu.Lock()
	defer jq.mu.Unlock()
	if jq.items.Len() == 0 {
		return nil
	}
	item := heap.Pop(&jq.items).(*queueItem)
	return item.job
}

// Peek returns the highest priority job without removing it
func (jq *JobQueue) Peek() *job.Job {
	jq.mu.RLock()
	defer jq.mu.RUnlock()
	if jq.items.Len() == 0 {
		return nil
	}
	return jq.items[0].job
}

// Len returns the number of jobs in the queue
func (jq *JobQueue) Len() int {
	jq.mu.RLock()
	defer jq.mu.RUnlock()
	return jq.items.Len()
}

// IsEmpty checks if the queue is empty
func (jq *JobQueue) IsEmpty() bool {
	return jq.Len() == 0
}

// queueItem represents an item in the priority queue
type queueItem struct {
	job      *job.Job
	priority int
	index    int
}
