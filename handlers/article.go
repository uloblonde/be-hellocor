package handlers

import (
	"fmt"
	articledto "halocorona/dto/article"
	dto "halocorona/dto/result"
	"halocorona/models"
	"halocorona/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var path_file = "http://localhost:5000/uploads/"

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

	for i, m := range article {
		article[i].ThumbnailArticle = path_file + m.ThumbnailArticle
	}

	return c.JSON(http.StatusOK, article)
}

func (h *handlerArticle) DapatArticle(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	article, err := h.ArticleRepository.DapatArticle(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	article.ThumbnailArticle = path_file + article.ThumbnailArticle

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
		UserId:           id,
		ThumbnailArticle: dataFile,
		Description:      c.FormValue("description"),
	}

	validation := validator.New()
	err := validation.Struct(meminta)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	// UserId:           meminta.UserId,

	article := models.Article{
		Title:            meminta.Title,
		ThumbnailArticle: meminta.ThumbnailArticle,
		UserId:           meminta.UserId,
		Description:      meminta.Description,
	}

	data, err := h.ArticleRepository.MembuatArticle(article)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
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
		Title:            u.Title,
		ThumbnailArticle: u.ThumbnailArticle,
		Description:      u.Description,
	}
}
func convertDeleteArticle(u models.Article) articledto.ArticleDeleteResponse {
	return articledto.ArticleDeleteResponse{
		Id: u.Id,
	}
}
