package cmd

import (
	"fmt"

	"github.com/nskforward/playbook/conn"
)

func FileExists(c *conn.Conn, path string) bool {
	cmd := fmt.Sprintf("[ -f %s ] || echo false", path)
	output := c.Execute(cmd)
	return output != "false"
}
