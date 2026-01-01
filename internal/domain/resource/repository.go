package resource

// Repository defines data access interface for resources
type Repository interface {
	Create(resource *Resource) error
	GetByID(rid uint) (*Resource, error)
	GetByConfigFileID(cfid uint) ([]Resource, error)
	GetByType(resourceType ResourceType) ([]Resource, error)
	List() ([]Resource, error)
	Update(resource *Resource) error
	Delete(rid uint) error
}
