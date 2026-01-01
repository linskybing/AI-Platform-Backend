package resource

import (
	"gorm.io/datatypes"
	"time"
)

// ResourceType defines the type of Kubernetes resource
type ResourceType string

const (
	ResourcePod        ResourceType = "pod"        // Kubernetes Pod
	ResourceService    ResourceType = "service"    // Kubernetes Service
	ResourceDeployment ResourceType = "deployment" // Kubernetes Deployment
	ResourceConfigMap  ResourceType = "configmap"  // Kubernetes ConfigMap
	ResourceIngress    ResourceType = "ingress"    // Kubernetes Ingress
	ResourceJob        ResourceType = "job"        // Kubernetes Job
)

// Resource represents a Kubernetes resource configuration
type Resource struct {
	RID         uint           `gorm:"primaryKey;column:r_id"`
	CFID        uint           `gorm:"not null;column:cf_id"` // ConfigFile ID
	Type        ResourceType   `gorm:"type:resource_type;not null"`
	Name        string         `gorm:"size:50;not null"`
	ParsedYAML  datatypes.JSON `gorm:"type:jsonb;not null;"`
	Description *string        `gorm:"type:text"`
	CreatedAt   time.Time      `gorm:"column:create_at;autoCreateTime"`
}

// TableName specifies the database table name
func (Resource) TableName() string {
	return "resource_list"
}

// IsPod checks if resource is a Pod
func (r *Resource) IsPod() bool {
	return r.Type == ResourcePod
}

// IsJob checks if resource is a Job
func (r *Resource) IsJob() bool {
	return r.Type == ResourceJob
}

// IsDeployment checks if resource is a Deployment
func (r *Resource) IsDeployment() bool {
	return r.Type == ResourceDeployment
}

type ResourceSwagger struct {
	RID         uint                   `json:"r_id"`
	CFID        uint                   `json:"cf_id"`
	Type        string                 `json:"type"`
	Name        string                 `json:"name"`
	ParsedYAML  map[string]interface{} `json:"parsedYAML" swaggertype:"object"`
	Description *string                `json:"description"`
	CreateAt    time.Time              `json:"create_at"`
}
