package course

// Repository defines data access interface for course workloads
type Repository interface {
	Create(workload *CourseWorkload) error
	GetByID(id uint) (*CourseWorkload, error)
	GetByUserID(userID uint) ([]CourseWorkload, error)
	GetByProjectID(projectID uint) ([]CourseWorkload, error)
	GetByStatus(status CourseWorkloadStatus) ([]CourseWorkload, error)
	GetPending() ([]CourseWorkload, error)
	Update(workload *CourseWorkload) error
	UpdateStatus(id uint, status CourseWorkloadStatus) error
	Delete(id uint) error
}
