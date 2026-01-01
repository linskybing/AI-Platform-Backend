package repository

import (
	"errors"

	"github.com/linskybing/platform-go/internal/config/db"
	"github.com/linskybing/platform-go/internal/domain/resource"
	"gorm.io/gorm"
)

type ResourceRepo interface {
	CreateResource(resource *resource.Resource) error
	GetResourceByID(rid uint) (*resource.Resource, error)
	UpdateResource(resource *resource.Resource) error
	DeleteResource(rid uint) error
	ListResourcesByProjectID(pid uint) ([]resource.Resource, error)
	ListResourcesByConfigFileID(cfID uint) ([]resource.Resource, error)
	GetResourceByConfigFileIDAndName(cfID uint, name string) (*resource.Resource, error)
}

type DBResourceRepo struct{}

func (r *DBResourceRepo) CreateResource(resource *resource.Resource) error {
	return db.DB.Create(resource).Error
}

func (r *DBResourceRepo) GetResourceByID(rid uint) (*resource.Resource, error) {
	var resource resource.Resource
	err := db.DB.First(&resource, "r_id = ?", rid).Error
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

func (r *DBResourceRepo) UpdateResource(resource *resource.Resource) error {
	if resource.RID == 0 {
		return errors.New("resource RID is required")
	}
	return db.DB.Save(resource).Error
}

func (r *DBResourceRepo) DeleteResource(rid uint) error {
	return db.DB.Delete(&resource.Resource{}, "r_id = ?", rid).Error
}

func (r *DBResourceRepo) ListResourcesByProjectID(pid uint) ([]resource.Resource, error) {
	var resources []resource.Resource
	err := db.DB.
		Joins("JOIN config_files cf ON cf.cf_id = resources.cf_id").
		Where("cf.project_id = ?", pid).
		Find(&resources).Error
	return resources, err
}

func (r *DBResourceRepo) ListResourcesByConfigFileID(cfID uint) ([]resource.Resource, error) {
	var resources []resource.Resource
	err := db.DB.
		Where("cf_id = ?", cfID).
		Find(&resources).Error
	return resources, err
}

func (r *DBResourceRepo) GetResourceByConfigFileIDAndName(cfID uint, name string) (*resource.Resource, error) {
	var resource resource.Resource
	err := db.DB.
		Where("cf_id = ? AND name = ?", cfID, name).
		First(&resource).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &resource, nil
}
