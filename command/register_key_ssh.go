package command

import (
	"fmt"
	"strings"

	"github.com/nskforward/playbook/conn"
	"github.com/nskforward/playbook/util"
)

func RegisterKeySSH(c *conn.Conn, user, pubkey string) bool {
	if user == "" {
		util.Check(fmt.Errorf("user name cannot be empty"))
	}

	parts := strings.Split(pubkey, " ")

	if len(parts) < 2 {
		util.Check(fmt.Errorf("bad pubkey format for user: %s", user))
	}

	if parts[0] != "ssh-rsa" {
		util.Check(fmt.Errorf("unknown pubkey ssh algo for user: %s", user))
	}

	search := strings.Join(parts[:2], " ")

	sshDir := fmt.Sprintf("/home/%s/.ssh", user)

	if !DirExists(c, sshDir) {
		DirCreate(c, true, sshDir)
		Chown(c, false, user, user, sshDir)
		Chmod(c, false, sshDir, util.NewPerm(7, 5, 0))
	}

	authorizedKeys := fmt.Sprintf("%s/authorized_keys", sshDir)

	if !FileExists(c, authorizedKeys) {
		FileCreateEmpty(c, authorizedKeys)
		Chown(c, false, user, user, authorizedKeys)
		Chmod(c, false, authorizedKeys, util.NewPerm(6, 0, 0))
	}

	if FileContains(c, authorizedKeys, search) {
		return false
	}

	FileWrite(c, true, authorizedKeys, pubkey)
	return true
}
