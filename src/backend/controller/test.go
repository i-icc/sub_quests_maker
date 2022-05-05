package controller

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

func GetTest() echo.HandlerFunc {
  return func(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, World!")
  }
}
