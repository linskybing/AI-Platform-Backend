package monitor

import (
	"context"
	"time"
)

// ResourceMetrics represents current resource usage
type ResourceMetrics struct {
	TotalCPU     float64
	UsedCPU      float64
	TotalMemory  int64
	UsedMemory   int64
	TotalGPU     int
	UsedGPU      int
	AvailableMPS int
	Timestamp    time.Time
	NodeName     string
	PodCount     int
	JobCount     int
	CourseCount  int
}

// CPUUsagePercent returns CPU usage percentage
func (r *ResourceMetrics) CPUUsagePercent() float64 {
	if r.TotalCPU == 0 {
		return 0
	}
	return (r.UsedCPU / r.TotalCPU) * 100
}

// MemoryUsagePercent returns memory usage percentage
func (r *ResourceMetrics) MemoryUsagePercent() float64 {
	if r.TotalMemory == 0 {
		return 0
	}
	return (float64(r.UsedMemory) / float64(r.TotalMemory)) * 100
}

// GPUUsagePercent returns GPU usage percentage
func (r *ResourceMetrics) GPUUsagePercent() float64 {
	if r.TotalGPU == 0 {
		return 0
	}
	return (float64(r.UsedGPU) / float64(r.TotalGPU)) * 100
}

// HasAvailableResources checks if there are available resources
func (r *ResourceMetrics) HasAvailableResources() bool {
	return r.UsedCPU < r.TotalCPU &&
		r.UsedMemory < r.TotalMemory &&
		(r.TotalGPU == 0 || r.UsedGPU < r.TotalGPU)
}

// Monitor defines the interface for resource monitoring
type Monitor interface {
	GetMetrics(ctx context.Context) (*ResourceMetrics, error)
	DetectResourceShortage(ctx context.Context) (bool, error)
	GetNodeMetrics(ctx context.Context, nodeName string) (*ResourceMetrics, error)
	Start(ctx context.Context) error
	Stop() error
}
