package mpi

import (
	"context"
	"fmt"
	"strings"
)

// MPIJobSpec defines specifications for an MPI job
type MPIJobSpec struct {
	JobID      uint
	Replicas   int
	Image      string
	Command    []string
	Args       []string
	WorkingDir string
	Env        map[string]string
	MPISlots   int // Slots per replica
}

// MPIExecutor handles MPI job execution
type MPIExecutor interface {
	Execute(ctx context.Context, spec *MPIJobSpec) error
	GenerateHostfile(ctx context.Context, spec *MPIJobSpec) (string, error)
	GetWorkerPods(ctx context.Context, jobID uint) ([]string, error)
	Terminate(ctx context.Context, jobID uint) error
}

// HostfileEntry represents an entry in MPI hostfile
type HostfileEntry struct {
	Hostname string
	Slots    int
}

// String returns hostfile entry format
func (h *HostfileEntry) String() string {
	return fmt.Sprintf("%s slots=%d", h.Hostname, h.Slots)
}

// GenerateHostfileContent creates hostfile content from entries
func GenerateHostfileContent(entries []HostfileEntry) string {
	var lines []string
	for _, entry := range entries {
		lines = append(lines, entry.String())
	}
	return strings.Join(lines, "\n")
}
