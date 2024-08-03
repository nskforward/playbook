package scenario

import (
	"fmt"
	"strings"

	"github.com/nskforward/playbook/cmd"
	"github.com/nskforward/playbook/conn"
	"github.com/nskforward/playbook/util"
)

func UserAddSSHKey(c *conn.Conn, user, pubkey string) {
	fmt.Println("# USER ADD SSH KEY", user)

	parts := strings.Split(pubkey, " ")
	if len(parts) < 2 {
		util.Check(fmt.Errorf("bad pubkey format for user: %s", user))
	}
	if parts[0] != "ssh-rsa" {
		util.Check(fmt.Errorf("unknown pubkey ssh algo for user: %s", user))
	}
	search := strings.Join(parts[:2], " ")

	if user == "" {
		util.Check(fmt.Errorf("user name cannot be empty"))
	}

	sshDir := fmt.Sprintf("/home/%s/.ssh", user)

	fmt.Println("--> check if ssh dir already exists")
	if cmd.DirExists(c, sshDir) {
		fmt.Println("<-- exists")
	} else {
		fmt.Println("<-- does not exist")
		fmt.Println("--> creating dir")
		cmd.DirMake(c, true, sshDir)
		fmt.Println("<-- ok")
		fmt.Println("--> change owner")
		cmd.Chown(c, false, user, user, sshDir)
		fmt.Println("<-- ok")
		fmt.Println("--> change permissions")
		cmd.Chmod(c, false, sshDir, util.NewPerm(7, 5, 0))
		fmt.Println("<-- ok")
	}

	authorizedKeys := fmt.Sprintf("%s/authorized_keys", sshDir)

	fmt.Println("--> check if authorized_keys file exists")
	if cmd.FileExists(c, authorizedKeys) {
		fmt.Println("<-- exists")

		fmt.Println("--> check if key already added")
		if cmd.FileContains(c, authorizedKeys, search) {
			fmt.Println("<-- contains")
			return
		} else {
			fmt.Println("<-- does not contain")

			fmt.Println("--> adding")
			cmd.FileWrite(c, true, authorizedKeys, pubkey)
			fmt.Println("<-- ok")
		}

	} else {
		fmt.Println("<-- does not exist")

		fmt.Println("--> creating file")
		cmd.FileWrite(c, false, authorizedKeys, pubkey)
		fmt.Println("<-- ok")
		fmt.Println("--> change owner")
		cmd.Chown(c, false, user, user, authorizedKeys)
		fmt.Println("<-- ok")
		fmt.Println("--> change permissions")
		cmd.Chmod(c, false, authorizedKeys, util.NewPerm(6, 0, 0))
		fmt.Println("<-- ok")
	}
}
