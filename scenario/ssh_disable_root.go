package scenario

import (
	"fmt"

	"github.com/nskforward/playbook/cmd"
	"github.com/nskforward/playbook/conn"
)

func SSHDisableRoot(c *conn.Conn) {
	fmt.Println("# DISABLE ROOT SSH LOGIN")

	fmt.Println("--> check if setting already exists")
	if cmd.FileExists(c, "/etc/ssh/ssh_config.d/disable_root_login.conf") {
		fmt.Println("<-- exists")
		return
	}
	fmt.Println("<-- does not exist")

	fmt.Println("--> adding")
	c.Execute(`echo "PermitRootLogin no" | sudo tee /etc/ssh/ssh_config.d/disable_root_login.conf`)
	fmt.Println("<-- ok")

	fmt.Println("--> restart sshd service")
	cmd.Systemctl(c, "sshd", cmd.Restart)
	fmt.Println("<-- ok")
}
