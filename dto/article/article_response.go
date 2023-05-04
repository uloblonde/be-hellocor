package articledto

import "halocorona/models"

type ArticleResponse struct {
	Id               int    `json:"id" `
	Title            string `json:"title" `
	User             models.User
	ThumbnailArticle string `json:"thumbnailArticle" `
	Description      string `json:"description" `
}
type ArticleDeleteResponse struct {
	Id int `json:"id" `
}
