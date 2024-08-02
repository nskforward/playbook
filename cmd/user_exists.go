package cmd

import (
	"fmt"

	"github.com/nskforward/playbook/conn"
)

func UserExists(c *conn.Conn, user string) bool {
	output := c.Execute(fmt.Sprintf("id %s || echo error", user))
	return output != "error"
}
