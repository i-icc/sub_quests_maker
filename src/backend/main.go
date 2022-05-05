package main

import (
	"backend/route"
)

func main() {
	e := route.Init()
	e.Logger.Fatal(e.Start(":3000"))
}
