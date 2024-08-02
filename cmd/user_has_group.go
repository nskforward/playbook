package cmd

import "github.com/nskforward/playbook/conn"

func UserHasGroup(c *conn.Conn, user, group string) bool {
	groups := UserGroups(c, user)
	for _, g := range groups {
		if g == group {
			return true
		}
	}
	return false
}
