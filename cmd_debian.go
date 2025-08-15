package playbook

import (
	"bytes"
	"fmt"

	"golang.org/x/crypto/ssh"
)

type DebianCommand struct {
	client *ssh.Client
}

func NewDebianCommand(client *ssh.Client) Command {
	return &DebianCommand{client}
}

func (cmd *DebianCommand) AddSSHKey(user, publicKey string) {
	sshDir := fmt.Sprintf("/home/%s/.ssh", user)
	if cmd.CreateDir(sshDir) {
		fmt.Println("+ .ssh dir created:", sshDir)
	} else {
		fmt.Println("+ .ssh dir already exists:", sshDir)
	}

	authorizedKeys := fmt.Sprintf("%s/authorized_keys", sshDir)

	if cmd.FileExist(authorizedKeys) {
		fmt.Println("+ file already exists:", authorizedKeys)
	} else {
		cmd.AppendToFile(authorizedKeys, publicKey)
		fmt.Println("+ ssh key registered:", authorizedKeys)
	}

	cmd.ChangePerm("700", sshDir)
	cmd.ChangePerm("600", authorizedKeys)
	cmd.ChangeOwner(fmt.Sprintf("%s:%s", user, user), sshDir)
	fmt.Println("+ added permissions")
}

func (cmd *DebianCommand) AddUserToSudo(user string) {
	cmd.AppendToFile(fmt.Sprintf("/etc/sudoers.d/%s", user), fmt.Sprintf("%s ALL=(ALL) NOPASSWD: ALL", user))
}

func (cmd *DebianCommand) ChangePassword(user string, password string) {
	command := fmt.Sprintf("sudo sh -c \"echo '%s' | passwd %s --stdin\"", password, user)
	_, err := execute(cmd.client, command, true)
	Catch(err, "cannot change password")
}

func (cmd *DebianCommand) CreateUser(user string) bool {
	if cmd.UserExists(user) {
		return false
	}
	_, err := execute(cmd.client, fmt.Sprintf("sudo adduser %s", user), true)
	Catch(err, "cannot create user")
	return true
}

func (cmd *DebianCommand) UserExists(user string) bool {
	output, err := execute(cmd.client, fmt.Sprintf("id %s || echo false", user), true)
	Catch(err, "cannot get user id")
	return !bytes.Equal(output, []byte("false"))
}

func (cmd *DebianCommand) CreateFile(path string) bool {
	if cmd.FileExist(path) {
		return false
	}
	_, err := execute(cmd.client, fmt.Sprintf("sudo touch %s", path), true)
	Catch(err, "cannot create empty file")
	return true
}

func (cmd *DebianCommand) FileExist(path string) bool {
	command := fmt.Sprintf("sudo sh -c \"[ -f %s ] || echo false\"", path)
	output, err := execute(cmd.client, command, true)
	Catch(err, "cannot check file location")
	return !bytes.Equal(output, []byte("false"))
}

func (cmd *DebianCommand) CreateDir(path string) bool {
	if cmd.DirExists(path) {
		return false
	}
	command := fmt.Sprintf("sudo mkdir -p %s", path)
	_, err := execute(cmd.client, command, true)
	Catch(err, "cannot check file location")
	return true
}

func (cmd *DebianCommand) DirExists(path string) bool {
	command := fmt.Sprintf("sudo sh -c \"[ -d %s ] || echo false\"", path)
	output, err := execute(cmd.client, command, true)
	Catch(err, "cannot check file location")
	return !bytes.Equal(output, []byte("false"))
}

func (cmd *DebianCommand) AppendToFile(path, data string) {
	if cmd.FileExist(path) {
		fmt.Println("+ file already exists:", path)
		return
	}
	cmd.CreateFile(path)
	command := fmt.Sprintf("sudo sh -c \"echo '%s' > %s\"", data, path)
	_, err := execute(cmd.client, command, true)
	Catch(err, "cannot append data to file")
}

func (cmd *DebianCommand) ChangeOwner(owner, path string) {
	command := fmt.Sprintf("sudo chown -R %s %s", owner, path)
	_, err := execute(cmd.client, command, true)
	Catch(err, "cannot change file owner")
}

func (cmd *DebianCommand) ChangePerm(perm, path string) {
	command := fmt.Sprintf("sudo chmod -R %s %s", perm, path)
	_, err := execute(cmd.client, command, true)
	Catch(err, "cannot change file permission")
}
