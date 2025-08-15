package main

import (
	"fmt"
	"os"

	"github.com/nskforward/playbook"
)

const (
	userToCreate = "sgorbachev"
)

var (
	hosts     = []string{"10.48.173.63"}
	publicKey = playbook.GetPublicKey(userToCreate)
)

func main() {
	privateKey := playbook.GetLocalFile(os.Getenv("SSH_PRIVATE"))

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
