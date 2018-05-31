package main

import "github.com/labstack/echo"

func addRoutes(e *echo.Echo) {
	e.GET("/", mainPage)
}
