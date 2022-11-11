package main

import (
	"echoinventory/config"
	"echoinventory/middleware"
	"echoinventory/routes"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db := config.DatabaseConnect()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	middleware.WebSecurityConfig(e)

	routes.NewRoute(db, e)

	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":5000"))
}
