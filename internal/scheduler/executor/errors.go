package executor

var (
	ErrUnsupportedJobType    = "unsupported job type"
	ErrJobExecutionFailed    = "job execution failed"
	ErrJobNotFound           = "job not found in Kubernetes"
	ErrLogRetrievalFailed    = "failed to retrieve job logs"
	ErrCancellationFailed    = "failed to cancel job"
	ErrInsufficientResources = "insufficient resources to execute job"
)
