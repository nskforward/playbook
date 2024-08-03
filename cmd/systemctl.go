package cmd

import (
	"fmt"

	"github.com/nskforward/playbook/conn"
	"github.com/nskforward/playbook/util"
)

type SystemctlAction string

var (
	Start   SystemctlAction = "start"
	Stop    SystemctlAction = "stop"
	Restart SystemctlAction = "restart"
	Enable  SystemctlAction = "enable"
	Disable SystemctlAction = "disable"
)

func Systemctl(c *conn.Conn, service string, action SystemctlAction) {
	output := c.Execute(fmt.Sprintf("systemctl --quiet %s %s", action, service))
	if output != "" {
		util.Check(fmt.Errorf("cmd.Systemctl failed: %s", output))
	}
}
