package routes

import (
	"halocorona/handlers"
	"halocorona/pkg/middleware"
	"halocorona/pkg/mysql"
	"halocorona/repositories"

	"github.com/labstack/echo/v4"
)

func ArticleRoute(e *echo.Group) {
	articleRepo := repositories.RepositoryArticle(mysql.DB)
	h := handlers.HandlerArticle(articleRepo)

	// e.GET("/artuser/:id", h.DapatCatId)
	e.GET("/CariArticle", h.CariArticle)
	e.GET("/article/:id", h.DapatArticle)
	e.POST("/buatarticle/:id", middleware.UploadFile(h.MembuatArticle))
	e.PATCH("/article/:id", middleware.Auth(middleware.UploadFile(h.UpdateArticle)))
	e.DELETE("/article/:id", middleware.Auth(h.HapusArticle))
}
