package cmd

import (
	"fmt"

	"github.com/nskforward/playbook/conn"
)

func FileMove(c *conn.Conn, src, dst string) {
	c.Execute(fmt.Sprintf("mv %s %s", src, dst))
}
