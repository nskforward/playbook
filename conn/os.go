package conn

import (
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

func GetOS(client *ssh.Client) OS {
	output := execute(client, "cat /etc/os-release | grep ^ID=")

	if output == "" {
		util.Check(fmt.Errorf("empty answer on parsing OS"))
	}

	options := strings.Split(output, "=")
	if len(options) != 2 {
		util.Check(fmt.Errorf("wrong answer on parsing OS info: %s", output))
	}

	return detectOS(strings.Trim(options[1], "\""))
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
