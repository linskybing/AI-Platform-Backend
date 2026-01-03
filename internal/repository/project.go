package repository

import (
	"github.com/linskybing/platform-go/internal/config/db"
	"github.com/linskybing/platform-go/internal/domain/project"
)

type ProjectRepo interface {
	GetProjectByID(id uint) (project.Project, error)
	GetGroupIDByProjectID(pID uint) (uint, error)
	CreateProject(p *project.Project) error
	UpdateProject(p *project.Project) error
	DeleteProject(id uint) error
	ListProjects() ([]project.Project, error)
	ListProjectsByGroup(id uint) ([]project.Project, error)
}

type DBProjectRepo struct{}

func (r *DBProjectRepo) GetProjectByID(id uint) (project.Project, error) {
	var project project.Project
	err := db.DB.First(&project, id).Error
	return project, err
}

func (r *DBProjectRepo) GetGroupIDByProjectID(pID uint) (uint, error) {
	var gID uint
	err := db.DB.Model(&project.Project{}).Select("g_id").Where("p_id = ?", pID).Scan(&gID).Error
	if err != nil {
		return 0, err
	}
	return gID, nil
}

func (r *DBProjectRepo) CreateProject(p *project.Project) error {
	if err := db.DB.Create(p).Error; err != nil {
		return err
	}

	// CRITICAL: Re-fetch from database to ensure we have the correct PID
	// In some cases, GORM's RETURNING clause doesn't populate p.PID correctly
	// This guarantees we return the actual database-generated ID
	var created project.Project
	if err := db.DB.Where("project_name = ? AND g_id = ?", p.ProjectName, p.GID).
		Order("create_at DESC"). // Use create_at to get the most recent one
		First(&created).Error; err != nil {
		return err
	}

	// Update the pointer with the correct PID from database
	p.PID = created.PID
	return nil
}

func (r *DBProjectRepo) UpdateProject(p *project.Project) error {
	return db.DB.Save(p).Error
}

func (r *DBProjectRepo) DeleteProject(id uint) error {
	return db.DB.Delete(&project.Project{}, id).Error
}

func (r *DBProjectRepo) ListProjects() ([]project.Project, error) {
	var projects []project.Project
	err := db.DB.Find(&projects).Error
	return projects, err
}

func (r *DBProjectRepo) ListProjectsByGroup(id uint) ([]project.Project, error) {
	var projects []project.Project
	if err := db.DB.Where("g_id = ?", id).Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}
