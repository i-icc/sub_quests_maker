package route

import (
	"backend/controller"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	api := e.Group("/api")
	{
		api.GET("/test", controller.GetTest())	
	}

	
	return e
}
