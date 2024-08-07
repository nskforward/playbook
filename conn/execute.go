package conn

import (
	"github.com/nskforward/playbook/util"
	"golang.org/x/crypto/ssh"
)

func execute(client *ssh.Client, command string) ([]byte, error) {
	session, err := client.NewSession()
	util.Check(err)
	defer session.Close()

	/*
		var b bytes.Buffer
		session.Stdout = &b
		//session.Stderr = &b
		err = session.Run(command)
		if err != nil {
			util.Check(fmt.Errorf("error: %w on command: %s, output: %s", err, command, b.String()))
		}
	*/

	return session.Output(command)

	/*
		if err != nil {
			util.Check(fmt.Errorf("error: %w on command: %s, output: %s", err, command, string(output)))
		}

		return strings.TrimSpace(string(output))
	*/
}
