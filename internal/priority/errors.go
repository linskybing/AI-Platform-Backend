package priority

// Error represents priority errors
type Error string

const (
	ErrJobNotRunning    = Error("job_not_running")
	ErrPreemptionFailed = Error("preemption_failed")
)

func (e Error) Error() string {
	return string(e)
}
