package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/ishansd94/go-reverse-proxy/internal/app"

	"os"
)

func main() {

	log.Info("[main]", "starting...")

	if err := app.Bootstrap(); err != nil {
		log.Error("[main]", "bootstrapping failed... ", err.Error())
		os.Exit(-1)
	}
}
