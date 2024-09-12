package command

import (
	"fmt"

	"github.com/nskforward/playbook/conn"
	"github.com/nskforward/playbook/util"
)

func UserAddGroup(c *conn.Conn, user, group string) bool {
	if UserHasGroup(c, user, group) {
		return false
	}
	output := c.Execute(fmt.Sprintf("usermod -a -G %s %s", group, user))
	if output != "" {
		util.Check(fmt.Errorf("cmd.UserAddGroup failed: %s", output))
	}
	return true
}
