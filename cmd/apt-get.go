package cmd

import (
	"strings"

	"github.com/nskforward/playbook/conn"
)

type AptGetAction string

var (
	Install AptGetAction = "install"
	Update  AptGetAction = "update"
)

func AptGet(c *conn.Conn, action AptGetAction, args ...string) {
	items := []string{"apt-get", "--quiet", "--yes", string(action)}
	items = append(items, args...)
	command := strings.Join(items, " ")
	c.Execute(command)
}
