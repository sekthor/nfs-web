package nfs

import (
	"log"
	"os"
	"regexp"
	"strings"
)

func WriteConfig(shares *[]NfsShare, configFile string) error {

	content := MarshalConfig(shares)

	f, err := os.Create(configFile)
	if err != nil {
		log.Fatal("could not open exports file")
	}
	defer f.Close()

	f.Write([]byte(content))
	return nil
}

// unmarshals the exports file and appends any found NFS shares to a
// given slice
func ReadConfig(shares *[]NfsShare, configFile string) error {
	data, err := os.ReadFile(configFile)

	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {

		// ignore line if it is a commment
		if ok, _ := regexp.MatchString(line, "^\\s*#"); ok || err != nil {
			continue
		}

		var share NfsShare
		UnmarshalShare(line, &share)
		*shares = append(*shares, share)
	}
	return nil
}

func ReadExportsFile(configFile string) (string, error) {
	data, err := os.ReadFile(configFile)

	if err != nil {
		return "", err
	}

	return string(data), nil
}
