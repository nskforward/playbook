package command

import (
	"fmt"

	"github.com/nskforward/playbook/conn"
)

func DirExists(c *conn.Conn, path string) bool {
	cmd := fmt.Sprintf("[ -d %s ] || echo false", path)
	output := c.Execute(cmd)
	return output != "false"
}

func FileExists(c *conn.Conn, path string) bool {
	cmd := fmt.Sprintf("[ -f %s ] || echo false", path)
	output := c.Execute(cmd)
	return output != "false"
}
