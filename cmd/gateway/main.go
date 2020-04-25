package main

import (
	"os"

	"github.com/sanity-io/litter"

	"github.com/ishansd94/gateway-proxy/internal/gateway"
	"github.com/ishansd94/gateway-proxy/pkg/log"
)

func main() {

	log.Info("gateway", "starting...")

	litter.Dump(struct{}{})

	if err := gateway.Run(); err != nil{
		log.Info("gateway", "exiting...")
		os.Exit(-1)
	}
}