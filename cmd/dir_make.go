package cmd

import (
	"bytes"
	"fmt"

	"github.com/nskforward/playbook/conn"
	"github.com/nskforward/playbook/util"
)

func DirMake(c *conn.Conn, createParent bool, path string) {
	var buf bytes.Buffer
	buf.WriteString("mkdir ")
	if createParent {
		buf.WriteString("-p ")
	}
	buf.WriteString(path)
	output := c.Execute(buf.String())
	if output != "" {
		util.Check(fmt.Errorf("cmd.DirMake failed: %s", output))
	}
}
