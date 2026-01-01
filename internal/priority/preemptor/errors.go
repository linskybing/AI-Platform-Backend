package preemptor

var (
	ErrNoJobsToPreempt                      = "no suitable jobs available for preemption"
	ErrPreemptionFailed                     = "failed to preempt job"
	ErrCheckpointFailed                     = "failed to create checkpoint"
	ErrJobNotPreemptible                    = "job is not preemptible"
	ErrInsufficientResourcesAfterPreemption = "insufficient resources even after preemption"
	ErrGracePeriodExpired                   = "grace period expired, forcing termination"
)
