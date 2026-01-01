package repository

import (
	"github.com/linskybing/platform-go/internal/config/db"
	"github.com/linskybing/platform-go/internal/domain/job"
)

type JobRepo interface {
	Create(job *job.Job) error
	FindAll() ([]job.Job, error)
	FindByID(id uint) (*job.Job, error)
	FindByUserID(userID uint) ([]job.Job, error)
	FindByNamespace(namespace string) ([]job.Job, error)
}

type DBJobRepo struct{}

func (r *DBJobRepo) Create(job *job.Job) error {
	return db.DB.Create(job).Error
}

func (r *DBJobRepo) FindAll() ([]job.Job, error) {
	var jobs []job.Job
	err := db.DB.Find(&jobs).Error
	return jobs, err
}

func (r *DBJobRepo) FindByID(id uint) (*job.Job, error) {
	var job job.Job
	err := db.DB.First(&job, id).Error
	return &job, err
}

func (r *DBJobRepo) FindByUserID(userID uint) ([]job.Job, error) {
	var jobs []job.Job
	err := db.DB.Where("user_id = ?", userID).Find(&jobs).Error
	return jobs, err
}

func (r *DBJobRepo) FindByNamespace(namespace string) ([]job.Job, error) {
	var jobs []job.Job
	err := db.DB.Where("namespace = ?", namespace).Find(&jobs).Error
	return jobs, err
}
