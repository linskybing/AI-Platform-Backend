package mps

const (
	MPSUnitsPerGPU  = 10  // 1 dedicated GPU = 10 MPS units
	DefaultMPSLimit = 100 // Default MPS thread percentage limit
	MaxMPSLimit     = 100 // Maximum MPS thread percentage (100%)
	MinMPSLimit     = 0   // Minimum MPS thread percentage
)

// ConvertGPUToMPS converts dedicated GPU count to MPS units
func ConvertGPUToMPS(gpuCount int) int {
	return gpuCount * MPSUnitsPerGPU
}

// ConvertMPSToGPU converts MPS units to equivalent GPU count
func ConvertMPSToGPU(mpsUnits int) int {
	gpus := mpsUnits / MPSUnitsPerGPU
	if mpsUnits%MPSUnitsPerGPU > 0 {
		gpus++ // Round up to ensure sufficient resources
	}
	return gpus
}

// ValidateMPSLimit validates if MPS limit is within acceptable range
func ValidateMPSLimit(limit int) bool {
	return limit >= MinMPSLimit && limit <= MaxMPSLimit
}

// CalculateMPSUsage calculates total MPS units from a list of requests
func CalculateMPSUsage(mpsRequests []int) int {
	total := 0
	for _, req := range mpsRequests {
		total += req
	}
	return total
}

// IsWithinQuota checks if requested MPS is within available quota
func IsWithinQuota(requested, available int) bool {
	return requested <= available
}
