package cmd

import (
	"fmt"

	"github.com/nskforward/playbook/conn"
)

func DirDelete(c *conn.Conn, path string) {
	c.Execute(fmt.Sprintf("rm -rf %s", path))
}
