package controller

import (
	"net/http"

	"backend/db"
	"backend/model"
	"github.com/labstack/echo/v4"
)

func GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		db := db.Connect()
		defer db.Close()

		var user []model.User
    	db.Raw("SELECT * FROM user").Scan(&user)

		return c.String(http.StatusOK, "sikatanai")
    	// return c.XML(http.StatusCreated, user)
  	}
}