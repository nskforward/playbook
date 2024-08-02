package scenario

import (
	"fmt"

	"github.com/nskforward/playbook/cmd"
	"github.com/nskforward/playbook/conn"
)

func UserCreate(c *conn.Conn, user string, system bool) {
	fmt.Println("# CREATE USER", user)

	fmt.Println("--> check if user exists:", user)
	if cmd.UserExists(c, user) {
		fmt.Println("<-- exists")
		return
	}

	fmt.Println("<-- does not exist")
	fmt.Println("--> add user:", user)

	cmd.UserAdd(c, user, system)
	fmt.Println("<-- ok")
}
