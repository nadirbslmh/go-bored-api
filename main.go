package main

import (
	"go-bored-api/database"
	"go-bored-api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDatabase()

	e := echo.New()

	routes.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
