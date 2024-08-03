package cmd

import (
	"fmt"
	"strings"

	"github.com/nskforward/playbook/conn"
	"github.com/nskforward/playbook/util"
)

type AptGetAction string

var (
	Install AptGetAction = "install"
	Update  AptGetAction = "update"
)

func AptGet(c *conn.Conn, action AptGetAction, args ...string) {
	items := []string{"apt-get", "--quiet", "--yes", string(action), "&&", "echo", "success"}
	items = append(items, args...)
	command := strings.Join(items, " ")
	output := c.Execute(command)
	if output != "" && !strings.HasSuffix(output, "success") {
		util.Check(fmt.Errorf("cmd.AptGet failed: %s", output))
	}
}
