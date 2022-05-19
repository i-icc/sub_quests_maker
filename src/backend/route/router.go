package route

import (
	"backend/controller"
	"net/http"
)

func Init() {


	api := e.Group("/api")
	{
		api.GET("/test", controller.GetTest())
		api.GET("/users", controller.GetUser())
	}
	auth := e.Group("/auth")
	{
		auth.GET("/login", controller.Login())
		// auth.GET("/logout", controller.Logout())
		// auth.GET("/signup", controller.Signout())
	}

	return e
}
