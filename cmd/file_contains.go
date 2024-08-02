package cmd

import (
	"fmt"

	"github.com/nskforward/playbook/conn"
)

func FileContains(c *conn.Conn, filePath, substring string) bool {
	output := c.Execute(fmt.Sprintf("sudo grep -q \"%s\" %s; [ $? -eq 0 ] || echo error", substring, filePath))
	return output != "error"
}
