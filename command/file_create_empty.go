package command

import (
	"fmt"

	"github.com/nskforward/playbook/conn"
)

func FileCreateEmpty(c *conn.Conn, path string) bool {
	if FileExists(c, path) {
		return false
	}
	c.Execute(fmt.Sprintf("touch %s", path))
	return true
}
