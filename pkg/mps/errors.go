package mps

var (
	ErrInvalidMPSLimit      = "MPS limit must be between 0 and 100"
	ErrInsufficientMPSQuota = "insufficient MPS quota"
	ErrInvalidMPSRequest    = "invalid MPS request"
	ErrConflictingGPUAndMPS = "cannot request both dedicated GPU and MPS"
	ErrNegativeMPSValue     = "MPS value cannot be negative"
)
