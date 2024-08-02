package cmd

import (
	"bytes"
	"fmt"

	"github.com/nskforward/playbook/conn"
	"github.com/nskforward/playbook/util"
)

func Chown(c *conn.Conn, recursive bool, user, group, path string) {
	var buf bytes.Buffer
	buf.WriteString("chown ")
	if recursive {
		buf.WriteString("-R ")
	}
	buf.WriteString(user)
	buf.WriteByte(':')
	buf.WriteString(group)
	buf.WriteByte(' ')
	buf.WriteString(path)

	output := c.Execute(buf.String())
	if output != "" {
		util.Check(fmt.Errorf("cmd.Chown failed: %s", output))
	}
}
