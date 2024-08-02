package scenario

import (
	"github.com/nskforward/playbook/cmd"
	"github.com/nskforward/playbook/conn"
)

func SSHDisableRoot(c *conn.Conn) {
	cmd.SSHConfig(c, "PermitRootLogin", "no")
}
