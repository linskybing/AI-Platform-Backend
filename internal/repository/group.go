package repository

import (
	"github.com/linskybing/platform-go/internal/config/db"
	"github.com/linskybing/platform-go/internal/domain/group"
)

type GroupRepo interface {
	GetAllGroups() ([]group.Group, error)
	GetGroupByID(id uint) (group.Group, error)
	CreateGroup(group *group.Group) error
	UpdateGroup(group *group.Group) error
	DeleteGroup(id uint) error
}

type DBGroupRepo struct{}

func NewGroupRepo() GroupRepo {
	return &DBGroupRepo{}
}
func (r *DBGroupRepo) GetAllGroups() ([]group.Group, error) {
	var groups []group.Group
	err := db.DB.Find(&groups).Error
	return groups, err
}

func (r *DBGroupRepo) GetGroupByID(id uint) (group.Group, error) {
	var group group.Group
	err := db.DB.First(&group, id).Error
	return group, err
}

func (r *DBGroupRepo) CreateGroup(group *group.Group) error {
	return db.DB.Create(group).Error
}

func (r *DBGroupRepo) UpdateGroup(group *group.Group) error {
	return db.DB.Save(group).Error
}

func (r *DBGroupRepo) DeleteGroup(id uint) error {
	return db.DB.Delete(&group.Group{}, id).Error
}
