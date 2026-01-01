package monitor

import "testing"

func TestCPUUsagePercent(t *testing.T) {
	m := &ResourceMetrics{TotalCPU: 8, UsedCPU: 2}
	if got := m.CPUUsagePercent(); got != 25 {
		t.Fatalf("expected 25%% CPU usage, got %f", got)
	}
}

func TestMemoryUsagePercent(t *testing.T) {
	m := &ResourceMetrics{TotalMemory: 1024, UsedMemory: 256}
	if got := m.MemoryUsagePercent(); got != 25 {
		t.Fatalf("expected 25%% memory usage, got %f", got)
	}
}

func TestGPUUsagePercentZeroTotal(t *testing.T) {
	m := &ResourceMetrics{TotalGPU: 0, UsedGPU: 1}
	if got := m.GPUUsagePercent(); got != 0 {
		t.Fatalf("expected 0%% GPU usage when total is zero, got %f", got)
	}
}

func TestHasAvailableResources(t *testing.T) {
	m := &ResourceMetrics{TotalCPU: 4, UsedCPU: 2, TotalMemory: 2048, UsedMemory: 1024, TotalGPU: 2, UsedGPU: 1}
	if !m.HasAvailableResources() {
		t.Fatalf("expected resources to be available")
	}
}
