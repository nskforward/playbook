package cmd

import (
	"fmt"

	"github.com/nskforward/playbook/conn"
	"github.com/nskforward/playbook/util"
)

func FileReplaceLine(c *conn.Conn, file, search, replace string) {
	command := fmt.Sprintf("sed -i '/%s/c\\%s' %s", search, replace, file)
	output := c.Execute(command)
	if output != "" {
		util.Check(fmt.Errorf("cmd.FileReplaceLine failed: %s", output))
	}
}
