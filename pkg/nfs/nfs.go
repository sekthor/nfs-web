package nfs

import (
	"errors"
	"fmt"
	"strings"
)

type NfsShare struct {
	Directory string `json:"directory"`
	Hosts     []Host `json:"hosts"`
}

type Host struct {
	Addr    string   `json:"addr"`
	Options []string `json:"options"`
}

func (h *Host) Marshal() string {

	if len(h.Options) == 0 {
		return h.Addr
	}

	return fmt.Sprintf("%s(%s)", h.Addr, strings.Join(h.Options, ","))
}

func UnmarshalShare(shareString string, share *NfsShare) error {
	elements := strings.Split(shareString, " ")

	if len(elements) < 2 {
		return ErrorInvalidShareFormat
	}

	// the first element is the source directory
	share.Directory = elements[0]

	// TODO: validate directory format
	// return ErrorInvalidDirectoryFormat

	for _, hostString := range elements[1:] {
		var host Host
		err := UnmarshalHost(hostString, &host)
		if err != nil {
			return err
		}
		share.Hosts = append(share.Hosts, host)
	}

	return nil
}

func UnmarshalHost(hostString string, host *Host) error {
	// if no options specified in brackets, we can just
	// copy the host
	if !strings.Contains(hostString, "(") {
		host.Addr = hostString
		return nil
	}

	host.Addr = hostString[:strings.IndexByte(hostString, '(')]
	options := string(hostString[strings.IndexByte(hostString, '('):])
	options = strings.Trim(options, "()")

	host.Options = strings.Split(options, ",")

	return nil
}

func (s *NfsShare) Marshal() string {
	var hosts []string

	for _, host := range s.Hosts {
		hosts = append(hosts, host.Marshal())
	}

	return s.Directory + " " + strings.Join(hosts, " ")
}

func MarshalConfig(shares *[]NfsShare) string {
	var content string

	for _, share := range *shares {
		content += share.Marshal() + "\n"
	}

	return content
}

func AppendShare(shares []NfsShare, share NfsShare) ([]NfsShare, error) {

	for _, s := range shares {
		if share.Directory == s.Directory {
			return shares, errors.New(
				"share with path " + share.Directory + " already exists")
		}
	}

	return append(shares, share), nil
}
