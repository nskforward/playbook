package scenario

import (
	"fmt"

	"github.com/nskforward/playbook/cmd"
	"github.com/nskforward/playbook/conn"
)

func UserSudo(c *conn.Conn, user string, askPass bool) {
	fmt.Println("# SUDO USER", user)

	fmt.Println("--> check if user already has sudo")
	if cmd.UserHasSudo(c, user) {
		fmt.Println("<-- has")
		return
	}

	fmt.Println("<-- no")
	fmt.Println("--> add sudo")

	cmd.UserAddSudo(c, user, askPass)
	fmt.Println("<-- ok")
}
