package main

import (
	"fmt"
	"halocorona/database"
	"halocorona/pkg/mysql"
	"halocorona/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	mysql.DatabaseInit()
	database.RunMigrations()
	routes.RouteInit(e.Group("/halo/v1"))

	fmt.Println("Server running on localhost:5000")
	e.Logger.Fatal(e.Start("localhost:5000"))
}
