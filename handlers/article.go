package handlers

import (
	"context"
	"fmt"
	articledto "halocorona/dto/article"
	dto "halocorona/dto/result"
	"halocorona/models"
	"halocorona/repositories"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerArticle struct {
	ArticleRepository repositories.ArticleRepository
}

func HandlerArticle(ArticleRepository repositories.ArticleRepository) *handlerArticle {
	return &handlerArticle{ArticleRepository}
}

func (h *handlerArticle) CariArticle(c echo.Context) error {
	article, err := h.ArticleRepository.CariArticle()

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, article)
}

func (h *handlerArticle) DapatArticle(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	article, err := h.ArticleRepository.DapatArticle(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: convertResArticle(article)})
}

func (h *handlerArticle) DapatCatId(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.ArticleRepository.DapatCatId(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: user})
}

func (h *handlerArticle) MembuatArticle(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	dataFile := c.Get("dataFile").(string)
	fmt.Println("this is data file", dataFile)

	meminta := articledto.CreatedArticleRequest{
		Title:            c.FormValue("title"),
		UserId:           int(id),
		ThumbnailArticle: dataFile,
		Description:      c.FormValue("description"),
	}

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, dataFile, uploader.UploadParams{Folder: "halocorona"})

	if err != nil {
		fmt.Println(err.Error())
	}

	validation := validator.New()
	err = validation.Struct(meminta)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	// UserId:           meminta.UserId,

	article := models.Article{
		Title:            meminta.Title,
		ThumbnailArticle: resp.SecureURL,
		UserId:           meminta.UserId,
		Description:      meminta.Description,
	}

	data, err := h.ArticleRepository.MembuatArticle(article)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	article, _ = h.ArticleRepository.DapatArticle(article.Id)

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: convertResArticle(data)})
}

func (h *handlerArticle) UpdateArticle(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)
	fmt.Println("this is data file", dataFile)

	meminta := articledto.CreatedArticleRequest{
		Title:            c.FormValue("title"),
		ThumbnailArticle: dataFile,
		Description:      c.FormValue("description"),
	}
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.ArticleRepository.DapatArticle(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	if meminta.Title != "" {
		user.Title = meminta.Title
	}
	if meminta.ThumbnailArticle != "" {
		user.ThumbnailArticle = meminta.ThumbnailArticle
	}
	if meminta.Description != "" {
		user.Description = meminta.Description
	}

	data, err := h.ArticleRepository.UpdateArticle(user, id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: convertResArticle(data)})
}

func (h *handlerArticle) HapusArticle(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	article, err := h.ArticleRepository.DapatArticle(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.ArticleRepository.HapusArticle(article, id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: convertDeleteArticle(data)})
}

func convertResArticle(u models.Article) articledto.ArticleResponse {
	return articledto.ArticleResponse{
		Id:               u.Id,
		UserId:           u.UserId,
		User:             u.User,
		Title:            u.Title,
		ThumbnailArticle: u.ThumbnailArticle,
		Description:      u.Description,
		CreatedAt:        u.CreatedAt,
		UpdatedAt:        u.UpdatedAt,
	}
}
func convertDeleteArticle(u models.Article) articledto.ArticleDeleteResponse {
	return articledto.ArticleDeleteResponse{
		Id: u.Id,
	}
}
