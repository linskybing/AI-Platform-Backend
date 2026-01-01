package job

import "time"

// CreateJobDTO represents a job creation request
type CreateJobDTO struct {
	Name        string       `json:"name" binding:"required"`
	Namespace   string       `json:"namespace" binding:"required"`
	Image       string       `json:"image" binding:"required"`
	Command     []string     `json:"command"`
	Args        []string     `json:"args"`
	Volumes     []VolumeSpec `json:"volumes"`
	GPUCount    int          `json:"gpu_count"`
	GPUType     string       `json:"gpu_type"`
	CPURequest  string       `json:"cpu_request"`
	Memory      string       `json:"memory"`
	Priority    string       `json:"priority"`
	Parallelism int32        `json:"parallelism"`
	Completions int32        `json:"completions"`
}

// JobSubmission represents a job submission request
type JobSubmission struct {
	Name        string       `json:"name"`
	Namespace   string       `json:"namespace"`
	Image       string       `json:"image"`
	Command     []string     `json:"command"`
	Args        []string     `json:"args"`
	Volumes     []VolumeSpec `json:"volumes"`
	GPUCount    int          `json:"gpu_count"`
	GPU         int          `json:"gpu"`
	GPUType     string       `json:"gpu_type"`
	CPURequest  string       `json:"cpu_request"`
	Memory      string       `json:"memory"`
	Priority    string       `json:"priority"`
	Parallelism int32        `json:"parallelism"`
	Completions int32        `json:"completions"`
}

// PVC represents a Persistent Volume Claim
type PVC struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Capacity  string `json:"capacity"`
	Status    string `json:"status"`
	PVCName   string `json:"pvc_name"`
	Size      string `json:"size"`
	IsGlobal  bool   `json:"is_global"`
}

// CreatePVCDTO represents a PVC creation request
type CreatePVCDTO struct {
	Name         string `json:"name" binding:"required"`
	Namespace    string `json:"namespace" binding:"required"`
	Capacity     string `json:"capacity" binding:"required"`
	StorageClass string `json:"storage_class"`
}

// ExpandPVCDTO represents a PVC expansion request
type ExpandPVCDTO struct {
	Name      string `json:"name" binding:"required"`
	Namespace string `json:"namespace" binding:"required"`
	Capacity  string `json:"capacity" binding:"required"`
}

// ExpandStorageInput represents storage expansion input
type ExpandStorageInput struct {
	StorageName  string `json:"storage_name" binding:"required"`
	NewSize      string `json:"new_size" binding:"required"`
	StorageClass string `json:"storage_class"`
}

// VolumeSpec defines a volume specification
type VolumeSpec struct {
	Name             string    `json:"name"`
	PVCName          string    `json:"pvc_name"`
	MountPath        string    `json:"mount_path"`
	Namespace        string    `json:"namespace"`
	StorageClassName string    `json:"storage_class_name"`
	Size             string    `json:"size"`
	Capacity         int       `json:"capacity"`
	Status           string    `json:"status"`
	AccessMode       string    `json:"access_mode"`
	ProjectID        uint      `json:"project_id"`
	ProjectName      string    `json:"project_name"`
	ID               uint      `json:"id"`
	CreatedAt        time.Time `json:"created_at"`
}

// ProjectStorageRequest represents a project storage request
type CreateProjectStorageRequest struct {
	ProjectID    uint   `json:"project_id"`
	ProjectName  string `json:"project_name"`
	Capacity     int    `json:"capacity"`
	StorageClass string `json:"storage_class"`
}

// ProjectPVCOutput represents a project PVC output
type ProjectPVCOutput struct {
	ID          uint      `json:"id"`
	ProjectID   uint      `json:"project_id"`
	ProjectName string    `json:"project_name"`
	Namespace   string    `json:"namespace"`
	Name        string    `json:"name"`
	Capacity    string    `json:"capacity"`
	Status      string    `json:"status"`
	AccessMode  string    `json:"access_mode"`
	CreatedAt   time.Time `json:"created_at"`
}
