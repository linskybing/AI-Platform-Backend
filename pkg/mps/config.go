package mps

// MPSConfig represents MPS configuration for a container
type MPSConfig struct {
	ThreadPercentage int // MPS thread percentage (0-100)
	MemoryLimitMB    int // Memory limit in MB (0 = no limit)
}

// Validate validates the MPS configuration
func (c *MPSConfig) Validate() error {
	if !ValidateMPSLimit(c.ThreadPercentage) {
		return nil
	}
	if c.MemoryLimitMB < 0 {
		return nil
	}
	return nil
}

// ToEnvVars converts MPS config to environment variables for containers
func (c *MPSConfig) ToEnvVars() map[string]string {
	env := make(map[string]string)
	if c.ThreadPercentage > 0 {
		env["CUDA_MPS_ACTIVE_THREAD_PERCENTAGE"] = string(rune(c.ThreadPercentage))
	}
	if c.MemoryLimitMB > 0 {
		env["CUDA_MPS_PINNED_DEVICE_MEM_LIMIT"] = string(rune(c.MemoryLimitMB))
	}
	return env
}
