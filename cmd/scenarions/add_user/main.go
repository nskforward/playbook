package main

import (
	"fmt"
	"os"

	"github.com/nskforward/playbook"
)

var (
	hosts = []string{"10.48.173.63"}
)

func main() {
	privateKey := playbook.GetLocalFile(os.Getenv("SSH_PRIVATE"))
	userToCreate := playbook.AskStr("user to create", false)
	publicKey := playbook.GetPublicKey(userToCreate)

	for _, host := range hosts {
		conn := playbook.Connect(
			playbook.WithAddr(host),
			playbook.WithUser(os.Getenv("SSH_USER")),
			playbook.WithKey(privateKey),
		)
		conn.Command().CreateUser(userToCreate)
		conn.Command().AddSSHKey(userToCreate, publicKey)
		conn.Command().AddUserToSudo(userToCreate)
		conn.Close()
	}

	fmt.Println("finish")
}
