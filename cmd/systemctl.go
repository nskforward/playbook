package cmd

import (
	"fmt"

	"github.com/nskforward/playbook/conn"
	"github.com/nskforward/playbook/util"
)

type Action string

var (
	Start   Action = "start"
	Stop    Action = "stop"
	Restart Action = "restart"
	Enable  Action = "enable"
	Disable Action = "disable"
)

func Systemctl(c *conn.Conn, service string, action Action) {
	output := c.Execute(fmt.Sprintf("systemctl --quiet %s %s", action, service))
	if output != "" {
		util.Check(fmt.Errorf("cmd.Systemctl failed: %s", output))
	}
}
