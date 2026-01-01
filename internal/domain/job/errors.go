package job

var (
	ErrJobNotFound           = "job not found"
	ErrInvalidJobName        = "invalid job name"
	ErrInvalidImage          = "invalid image"
	ErrInvalidMPIReplicas    = "MPI jobs require at least 2 replicas"
	ErrConflictingGPURequest = "cannot request both dedicated GPU and MPS"
	ErrInsufficientGPU       = "insufficient GPU quota"
	ErrInsufficientMPS       = "insufficient MPS quota"
	ErrJobNotPreemptible     = "job cannot be preempted"
	ErrJobAlreadyRunning     = "job is already running"
	ErrJobNotRunning         = "job is not running"
)
