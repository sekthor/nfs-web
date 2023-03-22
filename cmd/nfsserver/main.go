package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sekthor/nfs-server/pkg/nfs"
)

var config nfs.Config

func init() {

	err := config.LoadFromEnv()
	if err != nil {
		log.Fatal("could not load application config: ", err)
	}

	// make sure the exports file exists
	f, err := os.Create(config.ExportsFile)
	if err != nil {
		log.Fatal("could not open exports file: ", err)
	}
	f.Close()
}

func main() {
	router := gin.Default()

	router.GET("/config", GetConfig)
	router.POST("/shares", AddShare)
	router.DELETE("/shares", DeleteShare)

	err := router.Run(config.ServerHost + ":" + config.ServerPort)

	if err != nil {
		log.Fatal("could not start server")
	}
}
