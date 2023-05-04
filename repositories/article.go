package repositories

import (
	"halocorona/models"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	CariArticle() ([]models.Article, error)
	MembuatArticle(article models.Article) (models.Article, error)
	DapatArticle(Id int) (models.Article, error)
	UpdateArticle(article models.Article, Id int) (models.Article, error) //Update Article
	HapusArticle(article models.Article, Id int) (models.Article, error)
	DapatCatId(Id int) (models.User, error)
}

func RepositoryArticle(db *gorm.DB) *repo { //function Repository mengambil parameter berupa pointer ke gorm dan mengembalikan pointer ke repo
	return &repo{db}
}
func (r *repo) CariArticle() ([]models.Article, error) {
	var article []models.Article
	err := r.db.Preload("User").Find(&article).Error // Using Find method
	return article, err
}
func (r *repo) DapatArticle(Id int) (models.Article, error) {
	var article models.Article
	err := r.db.Preload("User").First(&article, Id).Error
	return article, err
}
func (r *repo) MembuatArticle(article models.Article) (models.Article, error) {
	err := r.db.Preload("User").Create(&article).Error
	return article, err
}
func (r *repo) UpdateArticle(article models.Article, Id int) (models.Article, error) {
	err := r.db.Preload("User").Model(&article).Updates(&article).Error
	return article, err
}

func (r *repo) HapusArticle(article models.Article, Id int) (models.Article, error) {
	err := r.db.Delete(&article).Error
	return article, err
}
func (r *repo) DapatCatId(Id int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, Id).Error
	return user, err
}
