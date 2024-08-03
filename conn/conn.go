package conn

import (
	"fmt"

	"github.com/nskforward/playbook/util"
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

func (c *Conn) Execute(cmd string) string {
	var err error
	var output []byte
	if c.sudo {
		output, err = execute(c.client, "sudo "+cmd)
	} else {
		output, err = execute(c.client, cmd)
	}
	if err != nil {
		util.Check(fmt.Errorf("failed command: '%s' > error: %w > output: %s", cmd, err, output))
	}
	return string(output)
}

func (c *Conn) OS() OS {
	return c.os
}
