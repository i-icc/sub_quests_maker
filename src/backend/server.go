package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
  app := iris.New()

  app.Handle("GET", "/api", func(ctx iris.Context) {
    ctx.JSON(iris.Map{"message": "ping"})
  })

  app.Handle("GET", "/", func(ctx iris.Context) {
    ctx.View("build/index.html")
})

  app.Listen(":3000")
}