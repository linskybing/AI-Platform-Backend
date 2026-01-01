package repository

import (
	"errors"
	"github.com/linskybing/platform-go/internal/domain/configfile"

	"github.com/linskybing/platform-go/internal/config/db"
)

type ConfigFileRepo interface {
	CreateConfigFile(cf *configfile.ConfigFile) error
	GetConfigFileByID(id uint) (*configfile.ConfigFile, error)
	UpdateConfigFile(cf *configfile.ConfigFile) error
	DeleteConfigFile(id uint) error
	ListConfigFiles() ([]configfile.ConfigFile, error)
	GetConfigFilesByProjectID(projectID uint) ([]configfile.ConfigFile, error)
}

type DBConfigFileRepo struct{}

func (r *DBConfigFileRepo) CreateConfigFile(cf *configfile.ConfigFile) error {
	return db.DB.Create(cf).Error
}

func (r *DBConfigFileRepo) GetConfigFileByID(id uint) (*configfile.ConfigFile, error) {
	var cf configfile.ConfigFile
	if err := db.DB.First(&cf, id).Error; err != nil {
		return nil, err
	}
	return &cf, nil
}

func (r *DBConfigFileRepo) UpdateConfigFile(cf *configfile.ConfigFile) error {
	if cf.CFID == 0 {
		return errors.New("missing ConfigFile ID")
	}
	return db.DB.Save(cf).Error
}

func (r *DBConfigFileRepo) DeleteConfigFile(id uint) error {
	return db.DB.Delete(&configfile.ConfigFile{}, id).Error
}

func (r *DBConfigFileRepo) ListConfigFiles() ([]configfile.ConfigFile, error) {
	var list []configfile.ConfigFile
	if err := db.DB.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *DBConfigFileRepo) GetConfigFilesByProjectID(projectID uint) ([]configfile.ConfigFile, error) {
	var files []configfile.ConfigFile
	if err := db.DB.Where("project_id = ?", projectID).Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}
