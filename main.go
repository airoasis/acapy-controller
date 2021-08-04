package main

import (
	"github.com/airoasis/controller/helper"
	"github.com/airoasis/controller/router"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

	helper.InitAcaPy()

	r := router.New()
	r.Logger.Fatal(r.Start(":8080"))

}

