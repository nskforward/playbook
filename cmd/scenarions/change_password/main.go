package main

import (
	"fmt"
	"os"

	"github.com/nskforward/playbook"
)

func main() {
	privateKey := playbook.GetLocalFile(os.Getenv("SSH_PRIVATE"))
	newPass := "test"
	userToChange := "test"
	cars := []string{}

	for _, car := range cars {
		conn := playbook.Connect(
			playbook.WithUser(os.Getenv("SSH_USER")),
			playbook.WithKey(privateKey),
			playbook.WithAddr(car),
		)
		conn.Command().ChangePassword(userToChange, newPass)
		fmt.Println("+", car)
		conn.Close()
	}
	fmt.Println("finish")
}
