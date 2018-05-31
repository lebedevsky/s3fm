package main

import (
	"net/http"
	"github.com/labstack/echo"
)

func mainPage(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}