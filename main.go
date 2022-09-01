package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/edgarSucre/moca/internal/controller"
	"github.com/edgarSucre/moca/internal/usecase"
	"github.com/edgarSucre/moca/pkg/httpserver"
)

func main() {

	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "02-01-2006 03:04:05 PM",
	})

	logger := log.WithFields(log.Fields{
		"layer": "usecase",
	})

	uc := usecase.New()
	handler := controller.New(uc, logger)
	server := httpserver.New(handler)

	log.Info("Listen on ", httpserver.DEFAULT_ADDRES)
	log.Fatal(server.Start())
}
