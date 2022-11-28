package repositories

import (
	"app/models"

	"gorm.io/gorm"
)

type DataRepository interface {
	ShowData(post models.Data) (models.Data, error)
}

func RepositoryData(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) ShowData(Post models.Data) (models.Data, error) {
	err := r.db.Create(&Post).Error

	return Post, err
}
