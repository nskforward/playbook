package cmd

import (
	"fmt"

	"github.com/nskforward/playbook/conn"
	"github.com/nskforward/playbook/util"
)

func GroupAddUser(c *conn.Conn, group, user string) {
	output := c.Execute(fmt.Sprintf("usermod -a -G %s %s", group, user))
	if output != "" {
		util.Check(fmt.Errorf("cmd.GroupAddUser failed: %s", output))
	}
}
