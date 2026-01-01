package course

import "time"

// CourseWorkloadStatus represents the status of a course workload (pod)
type CourseWorkloadStatus string

const (
	StatusPending   CourseWorkloadStatus = "Pending"   // Waiting to be scheduled
	StatusRunning   CourseWorkloadStatus = "Running"   // Currently running
	StatusSucceeded CourseWorkloadStatus = "Succeeded" // Completed successfully
	StatusFailed    CourseWorkloadStatus = "Failed"    // Failed to run
	StatusUnknown   CourseWorkloadStatus = "Unknown"   // Status unknown
)

// CoursePriority defines the priority level for course workloads
const (
	CoursePriority = 1000 // High priority for course pods
)

// CourseWorkload represents a high-priority course pod
type CourseWorkload struct {
	ID          uint                 `gorm:"primaryKey;column:id"`
	UserID      uint                 `gorm:"not null;column:user_id"`
	ProjectID   uint                 `gorm:"not null;column:project_id"`
	Name        string               `gorm:"size:100;not null"`
	Namespace   string               `gorm:"size:100;not null"`
	Image       string               `gorm:"size:255;not null"`
	Status      CourseWorkloadStatus `gorm:"size:50;default:'Pending'"`
	K8sPodName  string               `gorm:"size:100;not null;column:k8s_pod_name"`
	Priority    int                  `gorm:"default:1000"`
	ResourceCPU string               `gorm:"size:50;column:resource_cpu"`
	ResourceMem string               `gorm:"size:50;column:resource_mem"`
	ResourceGPU *int                 `gorm:"column:resource_gpu"`
	CreatedAt   time.Time            `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time            `gorm:"column:updated_at;autoUpdateTime"`
	StartedAt   *time.Time           `gorm:"column:started_at"`
}

// TableName specifies the database table name
func (CourseWorkload) TableName() string {
	return "course_workloads"
}

// RequiresGPU checks if this workload requires GPU resources
func (c *CourseWorkload) RequiresGPU() bool {
	return c.ResourceGPU != nil && *c.ResourceGPU > 0
}

// IsRunning checks if the workload is currently running
func (c *CourseWorkload) IsRunning() bool {
	return c.Status == StatusRunning
}

// IsPending checks if the workload is waiting to be scheduled
func (c *CourseWorkload) IsPending() bool {
	return c.Status == StatusPending
}
