package route

import (
	"log"
	"net/http"

	"backend/controller"
	"backend/oauth"

	_ "github.com/bradrydzewski/go.auth"
)

func Init() {
	// http.HandleFunc("/", auth.SecureUser(controller.GetTest))

	// http.HandleFunc("/api/test", controller.GetTest)
	http.HandleFunc("/api/usertest", controller.GetUserTest)

	http.HandleFunc("/api/cookie/set", controller.SetCookie)
	http.HandleFunc("/api/cookie/check", controller.ShowCookie)
	http.HandleFunc("/api/cookie/delete", controller.DeleteCookie)

	oauth.Init()
	o := oauth.Oauth{}
	o.SetUp()
	http.HandleFunc("/auth/login", o.Login)
	http.HandleFunc("/auth/callback", o.Callback)
	http.HandleFunc("/auth/logout", o.Logout)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
