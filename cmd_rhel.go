package playbook

import (
	"bytes"
	"fmt"

	"golang.org/x/crypto/ssh"
)

type RHELCommand struct {
	client *ssh.Client
}

func NewRHELCommand(client *ssh.Client) Command {
	return &RHELCommand{client}
}

func (cmd *RHELCommand) ChangePassword(user string, password string) {
	command := fmt.Sprintf("sudo sh -c \"echo '%s' | passwd %s --stdin\"", password, user)
	_, err := execute(cmd.client, command, true)
	Catch(err, "cannot change password")
}

func (cmd *RHELCommand) CreateUser(user string) bool {
	if cmd.UserExists(user) {
		return false
	}
	_, err := execute(cmd.client, fmt.Sprintf("sudo useradd %s", user), true)
	Catch(err, "cannot create user")
	return true
}

func (cmd *RHELCommand) AppendToFile(path string, data []byte) {
	if cmd.FileExist(path) {
		fmt.Println("+ file already exists:", path)
		return
	}
	cmd.CreateFile(path)
	command := fmt.Sprintf("sudo sh -c \"echo '%s' > %s\"", string(data), path)
	_, err := execute(cmd.client, command, true)
	Catch(err, "cannot append data to file")
}

func (cmd *RHELCommand) ChangeOwner(owner string, path string) {
	command := fmt.Sprintf("sudo chown -R %s %s", owner, path)
	_, err := execute(cmd.client, command, true)
	Catch(err, "cannot change file owner")
}

func (cmd *RHELCommand) ChangePerm(perm string, path string) {
	command := fmt.Sprintf("sudo chmod -R %s %s", perm, path)
	_, err := execute(cmd.client, command, true)
	Catch(err, "cannot change file permission")
}

func (cmd *RHELCommand) CreateDir(path string) bool {
	if cmd.DirExists(path) {
		return false
	}
	command := fmt.Sprintf("sudo mkdir -p %s", path)
	_, err := execute(cmd.client, command, true)
	Catch(err, "cannot check file location")
	return true
}

func (cmd *RHELCommand) CreateFile(path string) bool {
	if cmd.FileExist(path) {
		return false
	}
	_, err := execute(cmd.client, fmt.Sprintf("sudo touch %s", path), true)
	Catch(err, "cannot create empty file")
	return true
}

func (cmd *RHELCommand) DirExists(path string) bool {
	command := fmt.Sprintf("sudo sh -c \"[ -d %s ] || echo false\"", path)
	output, err := execute(cmd.client, command, true)
	Catch(err, "cannot check file location")
	return !bytes.Equal(output, []byte("false"))
}

func (cmd *RHELCommand) FileExist(path string) bool {
	command := fmt.Sprintf("sudo sh -c \"[ -f %s ] || echo false\"", path)
	output, err := execute(cmd.client, command, true)
	Catch(err, "cannot check file location")
	return !bytes.Equal(output, []byte("false"))
}

func (cmd *RHELCommand) UserExists(user string) bool {
	output, err := execute(cmd.client, fmt.Sprintf("id %s || echo false", user), true)
	Catch(err, "cannot get user id")
	return !bytes.Equal(output, []byte("false"))
}
