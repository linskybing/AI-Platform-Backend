package resource

import (
	"time"

	"gorm.io/datatypes"
)

type ResourceUpdateDTO struct {
	Type        *string         `json:"type,omitempty"`
	Name        *string         `json:"name,omitempty"`
	ParsedYAML  *datatypes.JSON `json:"parsed_yaml,omitempty"`
	Description *string         `json:"description,omitempty"`
}

type CreateJobDTO struct {
	Name        string   `json:"name" binding:"required"`
	Namespace   string   `json:"namespace" binding:"required"`
	Image       string   `json:"image" binding:"required"`
	Command     []string `json:"command" binding:"required"`
	Priority    string   `json:"priority"` // "high" or "low"
	GPUCount    int      `json:"gpu_count"`
	GPUType     string   `json:"gpu_type"` // "dedicated" or "shared"
	Parallelism int32    `json:"parallelism"`
	Completions int32    `json:"completions"`
	Volumes     []Volume `json:"volumes"`
}

type Volume struct {
	Name      string `json:"name"`
	PVCName   string `json:"pvc_name"`
	MountPath string `json:"mount_path"`
}

type CreatePVCDTO struct {
	Namespace        string `form:"namespace" binding:"required"`
	Name             string `form:"name" binding:"required"`
	StorageClassName string `form:"storageClassName" binding:"required"`
	Size             string `form:"size" binding:"required"`
}

type ExpandPVCDTO struct {
	Namespace string `form:"namespace" binding:"required"`
	Name      string `form:"name" binding:"required"`
	Size      string `form:"size" binding:"required"`
}

type ExpandStorageInput struct {
	NewSize string `json:"new_size" binding:"required" example:"1Ti"`
}

type PVC struct {
	Name      string `json:"name" example:"my-pvc"`
	Namespace string `json:"namespace" example:"default"`
	Status    string `json:"status" example:"Bound"`
	Size      string `json:"size" example:"1Gi"`
	IsGlobal  bool   `json:"isGlobal" example:"false"`
}

// CreateProjectStorageRequest defines the payload for creating project storage.
// 定義建立專案儲存空間的請求參數
type CreateProjectStorageRequest struct {
	ProjectID   uint   `json:"projectId" binding:"required"`
	ProjectName string `json:"projectName" binding:"required"`
	Capacity    int    `json:"capacity" binding:"required,min=1"` // In Gi
	Name        string `json:"name" binding:"required"`
}

// ProjectPVCOutput defines the response structure for listing storages.
type ProjectPVCOutput struct {
	ID          string    `json:"id"`          // The Project ID (string format to prevent frontend conversion issues)
	PVCName     string    `json:"pvcName"`     // The K8s PVC Name
	ProjectName string    `json:"projectName"` // Human readable name
	Namespace   string    `json:"namespace"`   // K8s Namespace
	Capacity    string    `json:"capacity"`    // e.g., "10Gi"
	Status      string    `json:"status"`      // e.g., "Bound"
	Role        string    `json:"role"`        // [NEW] User's role in the group (admin/manager/member)
	AccessMode  string    `json:"accessmode"`
	CreatedAt   time.Time `json:"createdAt"` // Creation timestamp
}

type StartFileBrowserDTO struct {
	Namespace string `json:"namespace" binding:"required"`
	PVCName   string `json:"pvc_name" binding:"required"`
}

type StopFileBrowserDTO struct {
	Namespace string `json:"namespace" binding:"required"`
	PVCName   string `json:"pvc_name" binding:"required"`
}
