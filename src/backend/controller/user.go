package controller

import (
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

		return c.String(200, "sikatanai")
    	// return c.Json(200, user)
  	}
}