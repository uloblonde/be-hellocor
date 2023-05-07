package repositories

import (
	"halocorona/models"

	"gorm.io/gorm"
)

type ResponseRepository interface {
	MembuatResponse(response models.Response) (models.Response, error)
	DapatResponse(Id uint) (models.Response, error)
	DapatResponseByConsul(Id uint) (models.Response, error)
}

func RepositoryResponse(db *gorm.DB) *repo {
	return &repo{db}
}

func (r *repo) MembuatResponse(response models.Response) (models.Response, error) {
	err := r.db.Create(&response).Error
	return response, err
}

func (r *repo) DapatResponse(Id uint) (models.Response, error) {
	var response models.Response
	err := r.db.Preload("User").Preload("Consulting").First(&response, Id).Error
	return response, err
}
func (r *repo) DapatResponseByConsul(Id uint) (models.Response, error) {
	var response models.Response
	err := r.db.Where("consul_id = ?", Id).Preload("User").Preload("Consulting").First(&response).Error
	return response, err
}
