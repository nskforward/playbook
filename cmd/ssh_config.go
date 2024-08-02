package cmd

import (
	"fmt"

	"github.com/nskforward/playbook/conn"
)

func SSHConfig(c *conn.Conn, param, value string) {
	FileReplaceLine(c, "/etc/ssh/sshd_config", param, fmt.Sprintf("%s %s", param, value))
}
