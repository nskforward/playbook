package cmd

import (
	"fmt"

	"github.com/nskforward/playbook/conn"
)

func FileDelete(c *conn.Conn, path string) {
	c.Execute(fmt.Sprintf("rm %s", path))
}
