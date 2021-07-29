package main

import (
	"example.com/controller/helper"
	"example.com/controller/router"
)

func main() {
	helper.InitAcaPy()

	r := router.New()
	r.Logger.Fatal(r.Start(":8080"))

}

