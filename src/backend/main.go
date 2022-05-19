package main

import (
	"backend/route"
	"log"
)

func main() {
	e := route.Init()
	log.Fatal(http.ListenAndServe(":3000", e))
}
