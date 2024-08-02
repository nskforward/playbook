package cmd

import (
	"fmt"
	"strings"

	"github.com/nskforward/playbook/conn"
	"github.com/nskforward/playbook/util"
)

func UserAddSudo(c *conn.Conn, user string, askPass bool) {
	if user == "" {
		util.Check(fmt.Errorf("user name cannot be empty"))
	}
	user = strings.ToLower(user)
	var command string
	if askPass {
		command = fmt.Sprintf("echo \"%s ALL=(ALL:ALL) ALL\" | sudo tee /etc/sudoers.d/%s", user, user)
	} else {
		command = fmt.Sprintf("echo \"%s ALL=(ALL:ALL) NOPASSWD: ALL\" | sudo tee /etc/sudoers.d/%s", user, user)
	}

	output := c.Execute(command)
	if output != command {
		util.Check(fmt.Errorf("cmd.UserSudo failed: %s", output))
	}
}

func UserHasSudo(c *conn.Conn, user string) bool {
	if user == "" {
		util.Check(fmt.Errorf("user name cannot be empty"))
	}
	user = strings.ToLower(user)
	file := fmt.Sprintf("/etc/sudoers.d/%s", user)
	return FileExists(c, file)
}
