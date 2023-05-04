package routes

import (
	"halocorona/handlers"
	"halocorona/pkg/middleware"
	"halocorona/pkg/mysql"
	"halocorona/repositories"

	"github.com/labstack/echo/v4"
)

func AuthRout(e *echo.Group) {
	authRepo := repositories.RepositoryAuth(mysql.DB)
	h := handlers.HandlerAuth(authRepo)

	e.POST("/register", h.Register)
	e.GET("/check-auth", middleware.Auth(h.CheckAuth))
	e.POST("/login", h.Login)
}
