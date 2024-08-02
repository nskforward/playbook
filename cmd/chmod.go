package cmd

import (
	"bytes"
	"fmt"

	"github.com/nskforward/playbook/conn"
	"github.com/nskforward/playbook/util"
)

func Chmod(c *conn.Conn, recursive bool, path string, perm util.Perm) {
	var buf bytes.Buffer
	buf.WriteString("chmod ")
	if recursive {
		buf.WriteString("-R ")
	}
	buf.WriteString(perm.String())
	buf.WriteByte(' ')
	buf.WriteString(path)

	output := c.Execute(buf.String())
	if output != "" {
		util.Check(fmt.Errorf("cmd.Chown failed: %s", output))
	}
}
