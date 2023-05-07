package routes

import (
	"halocorona/handlers"
	"halocorona/pkg/middleware"
	"halocorona/pkg/mysql"
	"halocorona/repositories"

	"github.com/labstack/echo/v4"
)

func ResponseRoutes(e *echo.Group) {
	responseRepository := repositories.RepositoryResponse(mysql.DB)
	h := handlers.HandlerResponse(responseRepository)

	e.POST("/response/:id", middleware.Auth(h.MembuatResponse))
	e.GET("/response/:id", h.DapatResponse)
	e.GET("/responseku/:id", h.DapatResponseByConsul)
}
