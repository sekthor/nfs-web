package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sekthor/nfs-server/pkg/nfs"
)

var config nfs.Config

func init() {

	config = nfs.ConfigFromEnv()
	defaults := nfs.Config{}
	defaults.LoadDefaults()
	config.Merge(&defaults)

	// make sure the exports file exists
	if _, err := os.Stat(config.ExportsFile); os.IsNotExist(err) {
		f, err := os.Create(config.ExportsFile)
		if err != nil {
			log.Fatal("could not open exports file: ", err)
		}
		f.Close()
	}
}

func main() {
	router := gin.Default()

	router.GET("/config", GetConfig)
	router.GET("/shares", GetShares)
	router.POST("/shares", AddShare)
	router.DELETE("/shares", DeleteShare)

	err := router.Run(config.ServerHost + ":" + config.ServerPort)

	if err != nil {
		log.Fatal("could not start server")
	}
}
