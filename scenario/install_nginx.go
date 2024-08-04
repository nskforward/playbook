package scenario

import (
	"fmt"

	"github.com/nskforward/playbook/cmd"
	"github.com/nskforward/playbook/conn"
)

func InstallNginx(c *conn.Conn) {
	fmt.Println("# INSTALL NGINX")

	cmd.Apt(c, cmd.Update)
	cmd.Apt(c, cmd.Install, "nginx")
	fmt.Println("<-- ok")
}
