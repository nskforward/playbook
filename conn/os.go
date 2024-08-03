package conn

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/nskforward/playbook/util"
	"golang.org/x/crypto/ssh"
)

type OS uint8

const (
	UNKNOWN OS = iota
	RHEL
	FEDORA
	DEBIAN
	SBERLINUX
	UBUNTU
)

func (os OS) String() string {
	switch os {
	case RHEL:
		return "rhel"

	case FEDORA:
		return "fedora"

	case DEBIAN:
		return "debian"

	case SBERLINUX:
		return "sberlinux"

	case UBUNTU:
		return "ubuntu"

	}
	return "unknown"
}

func getOS(client *ssh.Client) OS {
	output, err := execute(client, "cat /etc/os-release | grep ^ID=")
	if err != nil || len(output) == 0 {
		util.Check(fmt.Errorf("failed get os command: %w: %s", err, string(output)))
	}
	options := bytes.Split(output, []byte("="))
	if len(options) != 2 {
		util.Check(fmt.Errorf("wrong answer on parsing OS info: %s", string(output)))
	}
	return detectOS(string(bytes.Trim(options[1], "\"")))
}

func detectOS(option string) OS {
	switch strings.ToLower(option) {
	case "rhel":
		return RHEL
	case "fedora":
		return FEDORA
	case "debian":
		return DEBIAN
	case "sberlinux":
		return SBERLINUX
	case "ubuntu":
		return UBUNTU
	}

	util.Check(fmt.Errorf("unknown os family: %s", option))
	return UNKNOWN
}
