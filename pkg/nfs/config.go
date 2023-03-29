package nfs

import (
	"os"
	"strconv"
)

type Config struct {

	// file path for exports file
	ExportsFile string

	// port of the REST API
	ServerPort string

	// host of the REST API
	ServerHost string

	// the default path where directories for new nfs shares are written to
	DefaultSharePath string

	// the username for nfs share directory permissions
	NfsUser string

	// the name of the user group for share permissions
	NfsUserGroup string

	// manage the kernel
	ManageNfsServer bool
}

/**
 * Merge a config c2 into an existing config c1.
 * Copy all the values of c2's fileds into c1 where c1's fileds are of nil type.
 */
func (c1 *Config) Merge(c2 *Config) error {

	if c1.ExportsFile == "" {
		c1.ExportsFile = c2.ExportsFile
	}
	if c1.ServerPort == "" {
		c1.ServerPort = c2.ServerPort
	}
	if c1.ServerHost == "" {
		c1.ServerHost = c2.ServerHost
	}
	if c1.DefaultSharePath == "" {
		c1.DefaultSharePath = c2.DefaultSharePath
	}
	if c1.NfsUser == "" {
		c1.NfsUser = c2.NfsUser
	}
	if c1.NfsUserGroup == "" {
		c1.NfsUserGroup = c2.NfsUserGroup
	}
	c1.ManageNfsServer = c2.ManageNfsServer
	return nil
}

func ConfigFromEnv() Config {
	var c Config
	c.ExportsFile = os.Getenv("NFS_EXPORTS_FILE")
	c.ServerPort = os.Getenv("NFS_SERVER_PORT")
	c.ServerHost = os.Getenv("NFS_SERVER_HOST")
	c.DefaultSharePath = os.Getenv("NFS_DEFAULT_SHARE_PATH")
	c.NfsUser = os.Getenv("NFS_USER")
	c.NfsUserGroup = os.Getenv("NFS_USER_GROUP")
	c.ManageNfsServer, _ = strconv.ParseBool(os.Getenv("NFS_MANAGE_NFS_SERVER"))
	return c
}

func (c *Config) LoadDefaults() {
	c.ExportsFile = "/etc/exports"
	c.ServerPort = "8080"
	c.ServerHost = "0.0.0.0"
	c.DefaultSharePath = "/nfs"
	c.NfsUser = "nobody"
	c.NfsUserGroup = "nogroup"
	c.ManageNfsServer = false
}
