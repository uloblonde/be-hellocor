package articledto

type CreatedArticleRequest struct {
	Title            string `json:"title" form:"title" validate:"required"`
	UserId           int    `json:"userId" form:"userId" validate:"required"`
	ThumbnailArticle string `json:"thumbnailArticle" form:"thumbnailArticle" validate:"required"`
	Description      string `json:"description" form:"description" validate:"required"`
}
type UpdateArticleRequest struct {
	Title            string `json:"title" `
	ThumbnailArticle string `json:"thumbnailArticle" `
	Description      string `json:"description" `
}
