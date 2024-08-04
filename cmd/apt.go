package cmd

import (
	"strings"

	"github.com/nskforward/playbook/conn"
)

type AptAction string

var (
	Install    AptAction = "install"
	Remove     AptAction = "remove"
	Purge      AptAction = "purge"
	Update     AptAction = "update"
	Upgrade    AptAction = "upgrade"
	Autoremove AptAction = "autoremove"
)

func Apt(c *conn.Conn, action AptAction, args ...string) {
	items := []string{"apt-get", "--quiet", "--yes", string(action)}
	items = append(items, args...)
	command := strings.Join(items, " ")
	c.Execute(command)
}
