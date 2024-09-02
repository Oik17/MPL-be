package main

import (
	"net/http"

	"github.com/oik17/mpl-be/internal/controllers"
	"github.com/oik17/mpl-be/internal/database"

	"github.com/labstack/echo/v4"
)

func main() {
	database.Connect()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, Echo!"})
	})

	e.POST("/createTeam", controllers.CreateTeam)
	e.GET("/getAllTeams", controllers.GetAllTeams)
	e.Start(":8080")
}
