package models

import "time"

type Article struct {
	Id               int    `json:"id" gorm:"primary_key:auto_increment"`
	Title            string `json:"title" gorm:"type: varchar(255)" `
	User             User
	UserId           int    `json:"userId"`
	ThumbnailArticle string `json:"thumbnailArticle" gorm:"type: varchar(255)" `
	Description      string `json:"description" gorm:"type: varchar(255)" `
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
type ArticleResponse struct {
	Id               int    `json:"id" gorm:"primary_key:auto_increment"`
	Title            string `json:"title" gorm:"type: varchar(255)" `
	ThumbnailArticle string `json:"thumbnailArticle" gorm:"type: varchar(255)" `
	Description      string `json:"description" gorm:"type: varchar(255)" `
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (ArticleResponse) TableName() string {
	return "article_response"
}
