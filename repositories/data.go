package repositories

import (
	"app/models"

	"gorm.io/gorm"
)

type DataRepository interface {
	ShowData(DataUser string) (models.Data, error)
}

func RepositoryData(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) ShowData(DataUser string) (models.Data, error) {
	var data models.Data
	err := r.db.Where("data_user=?", DataUser).First(&data).Error

	return data, err
}
