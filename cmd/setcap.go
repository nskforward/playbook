package cmd

import "github.com/nskforward/playbook/conn"

func Setcap(c *conn.Conn, path string) {
	c.Execute("setcap 'cap_net_bind_service=+ep' " + path)
}
