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
	var output string
	if askPass {
		output = c.Execute(fmt.Sprintf("echo \"%s ALL=(ALL:ALL) ALL\" | sudo tee /etc/sudoers.d/%s", user, user))
	} else {
		output = c.Execute(fmt.Sprintf("echo \"%s ALL=(ALL:ALL) NOPASSWD: ALL\" | sudo tee /etc/sudoers.d/%s", user, user))
	}

	if output != "" {
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
