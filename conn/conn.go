package conn

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

type Conn struct {
	client *ssh.Client
	os     OS
	sudo   bool
}

func (c *Conn) Close() {
	c.client.Close()
	fmt.Println("# CONNECTION CLOSED")
}

func (c *Conn) Execute(cmd string) (output string) {
	if c.sudo {
		return execute(c.client, "sudo "+cmd)
	} else {
		return execute(c.client, cmd)
	}
}

func (c *Conn) OS() OS {
	return c.os
}
