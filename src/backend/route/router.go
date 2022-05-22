package route

import (
	"net/http"
	"log"

	"backend/controller"
	"backend/oauth"
	_ "github.com/bradrydzewski/go.auth"
)

func Init() {
	// http.HandleFunc("/", auth.SecureUser(controller.GetTest))

	// http.HandleFunc("/api/test", controller.GetTest)
	http.HandleFunc("/api/usertest", controller.GetUserTest)

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

