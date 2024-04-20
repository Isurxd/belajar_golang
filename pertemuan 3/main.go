package main

import (
	"net/http"
	"pertemuan3/routes"

	"github.com/labstack/echo/v4"
)

func main () {
	e :=echo.New()

	routes.StudentRoute(e)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello assalam")
	})

	e.Logger.Fatal(e.Start(":8080"))
}


// 