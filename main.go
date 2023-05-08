package main

import (
	"fmt"
	"halocorona/database"
	"halocorona/pkg/mysql"
	"halocorona/routes"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))

	e.Static("/uploads", "./uploads")

	var PORT = os.Getenv("PORT")

	mysql.DatabaseInit()
	database.RunMigrations()
	routes.RouteInit(e.Group("/halo/v1"))

	fmt.Println("server running localhost: " + PORT)
	e.Logger.Fatal(e.Start(":" + PORT))
}
