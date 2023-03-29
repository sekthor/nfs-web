package nfs

import "os/exec"

func IsNfsServerRunning() bool {
	return false
}

func IsNfsServerInstalled() bool {
	return false
}

func InstallNfsServer() error {
	return nil
}

func StartNfsServer() error {
	return nil
}

func RestartNfsServer() error {
	cmd := exec.Command("sudo", "apt", "restart", "nfs-kernel-server")
	err := cmd.Run()
	return err
}
