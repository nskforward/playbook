package cmd

import (
	"fmt"

	"github.com/nskforward/playbook/conn"
)

func FileCopy(c *conn.Conn, src, dst string) {
	c.Execute(fmt.Sprintf("cp %s %s", src, dst))
}
