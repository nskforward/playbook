package scenario

import (
	"fmt"

	"github.com/nskforward/playbook/cmd"
	"github.com/nskforward/playbook/conn"
)

func InstallFail2Ban(c *conn.Conn) {
	fmt.Println("# INSTALL FAIL2BAN")

	cmd.AptGet(c, cmd.Install, "fail2ban")
	fmt.Println("<-- ok")
}
