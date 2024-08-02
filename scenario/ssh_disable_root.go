package scenario

import (
	"fmt"

	"github.com/nskforward/playbook/cmd"
	"github.com/nskforward/playbook/conn"
)

func SSHDisableRoot(c *conn.Conn) {
	fmt.Println("# DISABLE ROOT SSH LOGIN")

	fmt.Println("--> edit ssh config")
	cmd.SSHConfig(c, "PermitRootLogin", "no")
	fmt.Println("<-- ok")

	fmt.Println("--> restart sshd service")
	cmd.Systemctl(c, "sshd", cmd.Restart)
	fmt.Println("<-- ok")
}
