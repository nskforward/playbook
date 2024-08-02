package cmd

import (
	"fmt"
	"strings"

	"github.com/nskforward/playbook/conn"
	"github.com/nskforward/playbook/util"
)

func UserGroups(c *conn.Conn, user string) []string {
	output := c.Execute(fmt.Sprintf("groups %s", user))
	parts := strings.Split(output, ":")
	if len(parts) == 0 {
		util.Check(fmt.Errorf("unknown output: %s", output))
	}
	return strings.Split(strings.TrimSpace(parts[1]), " ")
}
