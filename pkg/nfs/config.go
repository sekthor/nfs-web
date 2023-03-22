package nfs

import "os"

type Config struct {

	//
	ExportsFile string

	ServerPort string

	ServerHost string

	// the default path where directories for new nfs shares are written to
	DefaultSharePath string

	// the username for nfs share directory permissions
	NfsUser string

	// the name of the user group for share permissions
	NfsUserGroup string
}

func (c1 *Config) Merge(c2 *Config) error {

	// TODO: copy values of c2's fields into c1 where c1's fields are equal to ""

	return nil
}

func (c *Config) LoadFromEnv() error {
	c.ExportsFile = os.Getenv("NFS_EXPORTS_FILE")
	c.ServerPort = os.Getenv("NFS_SERVER_PORT")
	c.ServerHost = os.Getenv("NFS_SERVER_HOST")
	c.DefaultSharePath = os.Getenv("NFS_DEFAULT_SHARE_PATH")
	c.NfsUser = os.Getenv("NFS_USER")
	c.NfsUserGroup = os.Getenv("NFS_USER_GROUP")
	return nil
}

func (c *Config) LoadDefaults() {
	c.ExportsFile = "/etc/exports"
	c.ServerPort = "8080"
	c.ServerHost = "0.0.0.0"
	c.DefaultSharePath = "/nfs"
	c.NfsUser = "nobody"
	c.NfsUserGroup = "nogroup"
}
