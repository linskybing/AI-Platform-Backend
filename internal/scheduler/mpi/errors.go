package mpi

var (
	ErrInvalidReplicas          = "MPI jobs require at least 2 replicas"
	ErrHostfileGenerationFailed = "failed to generate MPI hostfile"
	ErrWorkerPodsNotReady       = "MPI worker pods are not ready"
	ErrMPIExecutionFailed       = "MPI job execution failed"
	ErrMPITerminationFailed     = "failed to terminate MPI job"
	ErrInvalidMPIConfig         = "invalid MPI configuration"
)
