package cmd

import (
	"fmt"

	"github.com/nskforward/playbook/conn"
)

type SystemctlAction string

var (
	Start   SystemctlAction = "start"
	Stop    SystemctlAction = "stop"
	Status  SystemctlAction = "status"
	Restart SystemctlAction = "restart"
	Enable  SystemctlAction = "enable"
	Disable SystemctlAction = "disable"
)

func Systemctl(c *conn.Conn, service string, action SystemctlAction) {
	c.Execute(fmt.Sprintf("systemctl %s %s", action, service))
}
