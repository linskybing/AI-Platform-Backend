package repository

import (
	"errors"

	"github.com/linskybing/platform-go/internal/config/db"
	"github.com/linskybing/platform-go/internal/domain/view"
	"gorm.io/gorm"
)

type ViewRepo interface {
	GetAllProjectGroupViews() ([]view.ProjectGroupView, error)
	GetProjectResourcesByGroupID(groupID uint) ([]view.ProjectResourceView, error)
	GetGroupResourcesByGroupID(groupID uint) ([]view.GroupResourceView, error)
	GetGroupIDByResourceID(rID uint) (uint, error)
	GetGroupIDByConfigFileID(cfID uint) (uint, error)
	IsSuperAdmin(uid uint) (bool, error)
	ListUsersByProjectID(projectID uint) ([]view.ProjectUserView, error)
	ListProjectsByUserID(userID uint) ([]view.ProjectUserView, error)
	GetUserRoleInGroup(uid uint, gid uint) (string, error)
}

type DBViewRepo struct{}

func (r *DBViewRepo) GetAllProjectGroupViews() ([]view.ProjectGroupView, error) {
	var results []view.ProjectGroupView
	err := db.DB.Find(&results).Error
	return results, err
}

func (r *DBViewRepo) GetProjectResourcesByGroupID(groupID uint) ([]view.ProjectResourceView, error) {
	var results []view.ProjectResourceView
	err := db.DB.Where("g_id = ?", groupID).Find(&results).Error
	return results, err
}

func (r *DBViewRepo) GetGroupResourcesByGroupID(groupID uint) ([]view.GroupResourceView, error) {
	var results []view.GroupResourceView
	err := db.DB.Where("g_id = ?", groupID).Find(&results).Error
	return results, err
}

func (r *DBViewRepo) GetGroupIDByResourceID(rID uint) (uint, error) {
	type result struct {
		GID uint `gorm:"column:g_id"`
	}

	var res result

	err := db.DB.Table("resources r").
		Select("p.g_id").
		Joins("JOIN config_files cf ON cf.cf_id = r.cf_id").
		Joins("JOIN project_list p ON cf.project_id = p.p_id").
		Where("r.r_id = ?", rID).
		Take(&res).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, gorm.ErrRecordNotFound
	}
	if err != nil {
		return 0, err
	}

	return res.GID, nil
}

func (r *DBViewRepo) GetGroupIDByConfigFileID(cfID uint) (uint, error) {
	type result struct {
		GID uint `gorm:"column:g_id"`
	}

	var res result

	err := db.DB.Table("config_files cf").
		Select("p.g_id").
		Joins("JOIN project_list p ON cf.project_id = p.p_id").
		Where("cf.cf_id = ?", cfID).
		Take(&res).Error

	if err != nil {
		return 0, err
	}

	return res.GID, nil
}

func (r *DBViewRepo) IsSuperAdmin(uid uint) (bool, error) {
	var view view.UserGroupView
	err := db.DB.
		Where("u_id = ? AND group_name = ? AND role = ?", uid, "super", "admin").
		First(&view).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *DBViewRepo) ListUsersByProjectID(projectID uint) ([]view.ProjectUserView, error) {
	var users []view.ProjectUserView
	err := db.DB.Where("p_id = ?", projectID).Find(&users).Error
	return users, err
}

func (r *DBViewRepo) ListProjectsByUserID(userID uint) ([]view.ProjectUserView, error) {
	var projects []view.ProjectUserView
	if err := db.DB.Where("u_id = ?", userID).Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

// GetUserRoleInGroup fetches the role of a user in a specific group.
func (r *DBViewRepo) GetUserRoleInGroup(uid uint, gid uint) (string, error) {
	roles := []string{}

	// Prefer "role" column
	err := db.DB.Table("user_group").
		Where("u_id = ? AND g_id = ?", uid, gid).
		Pluck("role", &roles).Error
	if err != nil {
		return "", err
	}
	if len(roles) > 0 {
		return roles[0], nil
	}

	// Fallback to legacy "role_level" column if present
	roles = []string{}
	err = db.DB.Table("user_group").
		Where("u_id = ? AND g_id = ?", uid, gid).
		Pluck("role_level", &roles).Error
	if err != nil {
		return "", err
	}
	if len(roles) > 0 {
		return roles[0], nil
	}

	return "", gorm.ErrRecordNotFound
}
