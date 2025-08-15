package playbook

import (
	"bytes"
	"fmt"

	"golang.org/x/crypto/ssh"
)

func execute(client *ssh.Client, cmd string, verbose bool) ([]byte, error) {
	if verbose {
		fmt.Println("-->", cmd)
	}
	session, err := client.NewSession()
	Catch(err, "cannot create a new ssh session")
	defer session.Close()
	output, err := session.Output(cmd)
	output = bytes.TrimSpace(output)
	if verbose {
		fmt.Println("<--", string(output))
	}
	return output, err
}
