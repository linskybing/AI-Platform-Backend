package repository

import (
	"github.com/linskybing/platform-go/internal/config/db"
	"github.com/linskybing/platform-go/internal/domain/group"
)

type UserGroupRepo interface {
	CreateUserGroup(userGroup *group.UserGroup) error
	UpdateUserGroup(userGroup *group.UserGroup) error
	DeleteUserGroup(uid, gid uint) error
	GetUserGroupsByUID(uid uint) ([]group.UserGroup, error)
	GetUserGroupsByGID(gid uint) ([]group.UserGroup, error)
	GetUserGroup(uid, gid uint) (group.UserGroup, error)
}

type DBUserGroupRepo struct{}

func (r *DBUserGroupRepo) CreateUserGroup(userGroup *group.UserGroup) error {
	return db.DB.Create(userGroup).Error
}

func (r *DBUserGroupRepo) UpdateUserGroup(userGroup *group.UserGroup) error {
	return db.DB.Save(userGroup).Error
}

func (r *DBUserGroupRepo) DeleteUserGroup(uid, gid uint) error {
	return db.DB.Where("u_id = ? AND g_id = ?", uid, gid).Delete(&group.UserGroup{}).Error
}

func (r *DBUserGroupRepo) GetUserGroupsByUID(uid uint) ([]group.UserGroup, error) {
	var userGroups []group.UserGroup
	err := db.DB.
		Where("u_id = ?", uid).
		Find(&userGroups).Error
	return userGroups, err
}

func (r *DBUserGroupRepo) GetUserGroupsByGID(gid uint) ([]group.UserGroup, error) {
	var userGroups []group.UserGroup
	err := db.DB.
		Where("g_id = ?", gid).
		Find(&userGroups).Error
	return userGroups, err
}

func (r *DBUserGroupRepo) GetUserGroup(uid, gid uint) (group.UserGroup, error) {
	var userGroup group.UserGroup
	err := db.DB.First(&userGroup, "u_id = ? AND g_id = ?", uid, gid).Error
	return userGroup, err
}
