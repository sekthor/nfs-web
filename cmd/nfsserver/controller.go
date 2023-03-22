package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sekthor/nfs-server/pkg/nfs"
)

func GetConfig(c *gin.Context) {
	content, err := nfs.ReadExportsFile(config.ExportsFile)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "could not read config file",
		})
	}

	c.Data(http.StatusOK, "text/plain", []byte(content))
}

func GetShares(c *gin.Context) {
	var shares []nfs.NfsShare
	err := nfs.ReadConfig(&shares, config.ExportsFile)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "could not read config file",
		})
	}

	c.JSON(http.StatusOK, shares)
}

func AddShare(c *gin.Context) {
	var shares []nfs.NfsShare
	var share nfs.NfsShare

	err := c.BindJSON(&share)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid format for nfs shares"})
		return
	}

	err = nfs.ReadConfig(&shares, config.ExportsFile)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not read existing config"})
		return
	}

	shares, err = nfs.AppendShare(shares, share)
	err = nfs.WriteConfig(&shares, config.ExportsFile)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not write configuration"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "share added"})
}

func DeleteShare(c *gin.Context) {
	var shares []nfs.NfsShare
	var share nfs.NfsShare

	err := c.BindJSON(&share)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid format for nfs share"})
		return
	}

	err = nfs.ReadConfig(&shares, config.ExportsFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not read exiting config"})
		return
	}

	for i, s := range shares {
		if s.Directory == share.Directory {
			shares = append(shares[:i], shares[i+1:]...)
			break
		}
	}

	err = nfs.WriteConfig(&shares, config.ExportsFile)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not write configuration"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"msg": "record deleted"})
}
