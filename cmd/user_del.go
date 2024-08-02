package cmd

import (
	"fmt"

	"github.com/nskforward/playbook/conn"
	"github.com/nskforward/playbook/util"
)

func UserDel(c *conn.Conn, user string) {
	var output string

	switch c.OS() {

	case conn.DEBIAN, conn.UBUNTU:
		output = c.Execute(fmt.Sprintf("deluser --remove-home %s", user))

	case conn.FEDORA, conn.RHEL, conn.SBERLINUX:
		output = c.Execute(fmt.Sprintf("userdel -r %s", user))
	}

	if output != "" {
		util.Check(fmt.Errorf("cmd.UserDel failed: %s", output))
	}
}
