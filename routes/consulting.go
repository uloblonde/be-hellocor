package routes

import (
	"halocorona/handlers"
	"halocorona/pkg/middleware"
	"halocorona/pkg/mysql"
	"halocorona/repositories"

	"github.com/labstack/echo/v4"
)

func ConsultingRoutes(e *echo.Group) {
	ConsultingRepository := repositories.RepositoryConsulting(mysql.DB)
	h := handlers.HandlerConsulting(ConsultingRepository)

	e.POST("/consulting", middleware.Auth(h.MembuatConsulting))
	e.GET("/consulting/:id", middleware.Auth(h.DapatConsulting))
	e.GET("/consultings/:id", h.CariConsultingKu)
	e.GET("/consultings", h.DapatConsul)
}
