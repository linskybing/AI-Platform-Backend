package repository

import (
	"github.com/linskybing/platform-go/internal/config/db"
	"github.com/linskybing/platform-go/internal/domain/user"
)

type UserRepo interface {
	GetAllUsers() ([]user.UserWithSuperAdmin, error)
	ListUsersPaging(page, limit int) ([]user.UserWithSuperAdmin, error)
	GetUserByID(id uint) (user.UserWithSuperAdmin, error)
	GetUsernameByID(id uint) (string, error)
	GetUserByUsername(username string) (user.User, error)
	GetUserRawByID(id uint) (user.User, error)
	SaveUser(user *user.User) error
	DeleteUser(id uint) error
}

type DBUserRepo struct{}

func (r *DBUserRepo) GetAllUsers() ([]user.UserWithSuperAdmin, error) {
	var users []user.UserWithSuperAdmin
	if err := db.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *DBUserRepo) GetUserByUsername(username string) (user.User, error) {
	var u user.User
	if err := db.DB.Where("username = ?", username).First(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func (r *DBUserRepo) ListUsersPaging(page, limit int) ([]user.UserWithSuperAdmin, error) {
	var users []user.UserWithSuperAdmin

	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	if err := db.DB.Offset(int(offset)).Limit(int(limit)).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *DBUserRepo) GetUserByID(id uint) (user.UserWithSuperAdmin, error) {
	var u user.UserWithSuperAdmin
	if err := db.DB.Table("users").Where("u_id = ?", id).First(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func (r *DBUserRepo) GetUsernameByID(id uint) (string, error) {
	var username string
	err := db.DB.Model(&user.User{}).Select("username").Where("u_id = ?", id).First(&username).Error
	if err != nil {
		return "", err
	}
	return username, nil
}

func (r *DBUserRepo) GetUserRawByID(id uint) (user.User, error) {
	var u user.User
	if err := db.DB.First(&u, id).Error; err != nil {
		return u, err
	}
	return u, nil
}

func (r *DBUserRepo) SaveUser(user *user.User) error {
	return db.DB.Save(user).Error
}

func (r *DBUserRepo) DeleteUser(id uint) error {
	return db.DB.Delete(&user.User{}, id).Error
}
