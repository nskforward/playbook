package command

import (
	"github.com/nskforward/playbook/conn"
)

func DisableRootLogin(c *conn.Conn) bool {
	if FileExists(c, "/etc/ssh/ssh_config.d/disable_root_login.conf") {
		return false
	}
	c.Execute(`echo "PermitRootLogin no" | sudo tee /etc/ssh/ssh_config.d/disable_root_login.conf`)
	Systemctl(c, "sshd", Restart)
	return true
}
