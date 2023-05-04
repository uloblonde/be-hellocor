package routes

import (
	"halocorona/handlers"
	"halocorona/pkg/mysql"
	"halocorona/repositories"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlersUser(userRepository)

	e.GET("/users", h.CariUser)
	e.GET("/profile", h.CariProfile)
	e.GET("/user/:id", h.DapatUser)
	e.POST("/user", h.MembuatUser)
	e.PATCH("/user/:id", h.UpdateUser)
	e.DELETE("/user/:id", h.HapusUser)
}
