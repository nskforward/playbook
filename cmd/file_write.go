package cmd

import (
	"fmt"

	"github.com/nskforward/playbook/conn"
	"github.com/nskforward/playbook/util"
)

func FileWrite(c *conn.Conn, append bool, file, body string) {
	var output string
	if append {
		output = c.Execute(fmt.Sprintf("sh -c 'echo \"%s\" >> %s'", body, file))
	} else {
		output = c.Execute(fmt.Sprintf("sh -c 'echo \"%s\" > %s'", body, file))
	}
	if output != "" {
		util.Check(fmt.Errorf("cmd.FileWrite failed: %s", output))
	}
}
