package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)


func CHandleGet(c echo.Context) error {
	return c.String(http.StatusOK, "Hellto, World!")
}
