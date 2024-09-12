package command

import "github.com/nskforward/playbook/conn"

func NetBindService(c *conn.Conn, path string) {
	c.Execute("setcap 'cap_net_bind_service=+ep' " + path)
}
