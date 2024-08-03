package conn

import (
	"bytes"
	"fmt"

	"github.com/nskforward/playbook/util"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type Conn struct {
	client *ssh.Client
	os     OS
	sudo   bool
	debug  bool
}

func (c *Conn) Close() {
	c.client.Close()
	fmt.Println("# CONNECTION CLOSED")
}

func (c *Conn) Execute(command string) string {
	var err error
	var output []byte
	var cmd string
	if c.sudo {
		command = "sudo " + command
	}

	if c.debug {
		fmt.Println("=======================================================================================================")
		fmt.Println(command)
		fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -")
	}

	output, err = execute(c.client, command)
	output = bytes.TrimSpace(output)

	if c.debug {
		if len(output) > 0 {
			fmt.Println(string(output))
		} else {
			fmt.Println("<empty>")
		}
		fmt.Println("=======================================================================================================")
	}

	if err != nil {
		if c.debug {
			util.Check(fmt.Errorf("failed command: '%s' > error: %w", cmd, err))
		} else {
			util.Check(fmt.Errorf("failed command: '%s' > error: %w > output: %s", cmd, err, output))
		}
	}

	return string(output)
}

func (c *Conn) OS() OS {
	return c.os
}

func (c *Conn) SFTP() *sftp.Client {
	sftp, err := sftp.NewClient(c.client)
	util.Check(err)
	return sftp
}
