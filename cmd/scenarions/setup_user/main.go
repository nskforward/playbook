package main

import (
	"fmt"
	"os"

	"github.com/nskforward/playbook"
)

func main() {
	pubKey := playbook.GetLocalFile(os.Getenv("SSH_PUBLIC"))

	conn := playbook.Connect()

	newUser := os.Getenv("SSH_USER")

	if conn.Command().CreateUser(newUser) {
		fmt.Println("+ user created:", newUser)
	} else {
		fmt.Println("+ user already exists:", newUser)
	}

	sshDir := fmt.Sprintf("/home/%s/.ssh", newUser)

	if conn.Command().CreateDir(sshDir) {
		fmt.Println("+ .ssh dir created:", sshDir)
	} else {
		fmt.Println("+ .ssh dir already exists:", sshDir)
	}

	authorizedKeys := fmt.Sprintf("%s/authorized_keys", sshDir)

	if conn.Command().FileExist(authorizedKeys) {
		fmt.Println("+ file already exists:", authorizedKeys)
	} else {
		conn.Command().AppendToFile(authorizedKeys, string(pubKey))
		fmt.Println("+ ssh key registered:", authorizedKeys)
	}

	conn.Command().ChangePerm("700", sshDir)
	conn.Command().ChangePerm("600", authorizedKeys)
	conn.Command().ChangeOwner(fmt.Sprintf("%s:%s", newUser, newUser), sshDir)
	fmt.Println("+ added permissions")

	conn.Command().AppendToFile(fmt.Sprintf("/etc/sudoers.d/%s", newUser), fmt.Sprintf("%s ALL=(ALL) NOPASSWD: ALL", newUser))
	fmt.Println("+ added to sudo")
}
