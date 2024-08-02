package cmd

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/nskforward/playbook/conn"
	"github.com/nskforward/playbook/util"
)

func UserAdd(c *conn.Conn, user string, system bool) {
	if user == "" {
		util.Check(fmt.Errorf("user name cannot be empty"))
	}
	user = strings.ToLower(user)
	var buf bytes.Buffer

	switch c.OS() {

	case conn.DEBIAN, conn.UBUNTU:
		buf.WriteString("adduser --quiet --gecos \"\" --disabled-password ")
		if system {
			buf.WriteString("--disabled-login --no-create-home ")
		}
		buf.WriteString(user)

	case conn.FEDORA, conn.RHEL, conn.SBERLINUX:
		buf.WriteString("useradd --user-group ")
		if system {
			buf.WriteString("--system ")
		} else {
			buf.WriteString("--create-home ")
		}
		buf.WriteString(user)
	}

	output := c.Execute(buf.String())
	if output != "" {
		util.Check(fmt.Errorf("cmd.UserAdd failed: %s", output))
	}
}
