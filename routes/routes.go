package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	UserRoutes(e)
	AuthRout(e)
	ArticleRoute(e)
	ConsultingRoutes(e)
	ResponseRoutes(e)
}
