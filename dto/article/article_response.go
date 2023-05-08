package articledto

import (
	"halocorona/models"
	"time"
)

type ArticleResponse struct {
	Id               int    `json:"id" `
	Title            string `json:"title" `
	UserId           int    `json:"user_id"`
	User             models.User
	ThumbnailArticle string `json:"thumbnailArticle" `
	Description      string `json:"description" `
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
type ArticleDeleteResponse struct {
	Id int `json:"id" `
}
