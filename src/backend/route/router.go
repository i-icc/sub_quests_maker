package route

import (
	"net/http"
	"log"

	"backend/controller"
	"github.com/bradrydzewski/go.auth"
)

func Init() {
	http.HandleFunc("/", auth.SecureFunc(controller.GetTest))

	http.HandleFunc("/api/test", controller.GetTest)
	http.HandleFunc("/api/usertest", controller.GetUserTest)

	http.HandleFunc("/auth/login", controller.Login)
	http.HandleFunc("/auth/logout", controller.Logout)

	err := http.ListenAndServe(":3000", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
