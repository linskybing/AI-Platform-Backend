package executor

import (
	"context"
	"errors"
	"testing"

	"github.com/linskybing/platform-go/internal/domain/job"
)

// MockExecutor for testing
type MockExecutor struct {
	executeErr  bool
	cancelErr   bool
	statusVal   job.JobStatus
	logs        string
	supportType bool
}

func (m *MockExecutor) Execute(ctx context.Context, j *job.Job) error {
	if m.executeErr {
		return errors.New("execute failed")
	}
	return nil
}

func (m *MockExecutor) Cancel(ctx context.Context, jobID uint) error {
	if m.cancelErr {
		return errors.New("cancel failed")
	}
	return nil
}

func (m *MockExecutor) GetStatus(ctx context.Context, jobID uint) (job.JobStatus, error) {
	return m.statusVal, nil
}

func (m *MockExecutor) GetLogs(ctx context.Context, jobID uint) (string, error) {
	return m.logs, nil
}

func (m *MockExecutor) SupportsType(jobType job.JobType) bool {
	return m.supportType
}

func TestNewExecutorRegistry(t *testing.T) {
	registry := NewExecutorRegistry()
	if registry == nil {
		t.Fatal("expected non-nil registry")
	}
	if registry.executors == nil {
		t.Fatal("expected executors map to be initialized")
	}
}

func TestRegisterAndGetExecutor(t *testing.T) {
	registry := NewExecutorRegistry()
	mockExec := &MockExecutor{supportType: true}

	registry.Register("test_type", mockExec)

	exec, exists := registry.GetExecutor("test_type")
	if !exists {
		t.Fatal("expected executor to exist")
	}
	if exec != mockExec {
		t.Fatal("expected same executor instance")
	}

	_, exists = registry.GetExecutor("unknown_type")
	if exists {
		t.Fatal("expected unknown executor to not exist")
	}
}

func TestRegisterOverwrite(t *testing.T) {
	registry := NewExecutorRegistry()
	exec1 := &MockExecutor{supportType: true}
	exec2 := &MockExecutor{supportType: false}

	registry.Register("type1", exec1)
	registry.Register("type1", exec2)

	exec, exists := registry.GetExecutor("type1")
	if !exists || exec != exec2 {
		t.Fatal("expected executor to be overwritten")
	}
}

func TestExecuteWithRegisteredExecutor(t *testing.T) {
	registry := NewExecutorRegistry()
	mockExec := &MockExecutor{supportType: true, executeErr: false}
	registry.Register("test", mockExec)

	j := &job.Job{ID: 1, JobType: "test"}
	err := registry.Execute(context.Background(), j)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestExecuteWithUnregisteredExecutor(t *testing.T) {
	registry := NewExecutorRegistry()

	j := &job.Job{ID: 1, JobType: "unknown"}
	err := registry.Execute(context.Background(), j)
	if err != ErrExecutorNotFound {
		t.Fatalf("expected ErrExecutorNotFound, got %v", err)
	}
}

func TestExecuteWithExecutorError(t *testing.T) {
	registry := NewExecutorRegistry()
	mockExec := &MockExecutor{supportType: true, executeErr: true}
	registry.Register("test", mockExec)

	j := &job.Job{ID: 1, JobType: "test"}
	err := registry.Execute(context.Background(), j)
	if err == nil {
		t.Fatal("expected error from executor")
	}
}

func TestMultipleExecutors(t *testing.T) {
	registry := NewExecutorRegistry()
	exec1 := &MockExecutor{supportType: true}
	exec2 := &MockExecutor{supportType: true}

	registry.Register("type1", exec1)
	registry.Register("type2", exec2)

	e1, exists1 := registry.GetExecutor("type1")
	e2, exists2 := registry.GetExecutor("type2")

	if !exists1 || !exists2 {
		t.Fatal("expected both executors to exist")
	}
	if e1 != exec1 || e2 != exec2 {
		t.Fatal("expected correct executor instances")
	}
}

func TestExecuteWithContextCancellation(t *testing.T) {
	registry := NewExecutorRegistry()
	mockExec := &MockExecutor{supportType: true}
	registry.Register("test", mockExec)

	j := &job.Job{ID: 1, JobType: "test"}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	// Should still execute even with cancelled context
	err := registry.Execute(ctx, j)
	if err != nil {
		t.Fatalf("expected no error with cancelled context, got %v", err)
	}
}
