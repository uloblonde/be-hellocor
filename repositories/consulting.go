package repositories

import (
	"halocorona/models"

	"gorm.io/gorm"
)

type ConsultingRepository interface {
	CariConsultingKu(Id int) ([]models.Consulting, error)
	DapatConsul() ([]models.Consulting, error)
	MembuatConsulting(article models.Consulting) (models.Consulting, error)
	DapatConsulting(Id uint) (models.Consulting, error)
	EditConsulting(article models.Consulting) (models.Consulting, error)
}

func RepositoryConsulting(db *gorm.DB) *repo { //function Repository mengambil parameter berupa pointer ke gorm dan mengembalikan pointer ke repo
	return &repo{db}
}
func (r *repo) CariConsultingKu(Id int) ([]models.Consulting, error) {
	var consul []models.Consulting
	err := r.db.Where("user_id = ?", Id).Preload("User").Find(&consul).Error
	return consul, err
}
func (r *repo) DapatConsulting(Id uint) (models.Consulting, error) {
	var consul models.Consulting
	err := r.db.Preload("User").First(&consul, Id).Error
	return consul, err
}
func (r *repo) MembuatConsulting(consul models.Consulting) (models.Consulting, error) {
	err := r.db.Create(&consul).Error
	return consul, err
}

func (r *repo) DapatConsul() ([]models.Consulting, error) {
	var consul []models.Consulting
	err := r.db.Preload("User").Find(&consul).Error
	return consul, err
}

func (r *repo) EditConsulting(consul models.Consulting) (models.Consulting, error) {
	err := r.db.Save(&consul).Error
	return consul, err
}
