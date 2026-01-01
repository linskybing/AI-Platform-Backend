package repository

import (
	"github.com/linskybing/platform-go/internal/config/db"
	"github.com/linskybing/platform-go/internal/domain/gpu"
)

type GPURequestRepo interface {
	Create(req *gpu.GPURequest) error
	Update(req *gpu.GPURequest) error
	GetByID(id uint) (gpu.GPURequest, error)
	ListByProjectID(projectID uint) ([]gpu.GPURequest, error)
	ListPending() ([]gpu.GPURequest, error)
}

type DBGPURequestRepo struct{}

func (r *DBGPURequestRepo) Create(req *gpu.GPURequest) error {
	return db.DB.Create(req).Error
}

func (r *DBGPURequestRepo) Update(req *gpu.GPURequest) error {
	return db.DB.Save(req).Error
}

func (r *DBGPURequestRepo) GetByID(id uint) (gpu.GPURequest, error) {
	var req gpu.GPURequest
	err := db.DB.First(&req, id).Error
	return req, err
}

func (r *DBGPURequestRepo) ListByProjectID(projectID uint) ([]gpu.GPURequest, error) {
	var reqs []gpu.GPURequest
	err := db.DB.Where("project_id = ?", projectID).Find(&reqs).Error
	return reqs, err
}

func (r *DBGPURequestRepo) ListPending() ([]gpu.GPURequest, error) {
	var reqs []gpu.GPURequest
	err := db.DB.Where("status = ?", gpu.GPURequestStatusPending).Find(&reqs).Error
	return reqs, err
}
