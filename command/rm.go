package command

import (
	"fmt"
	"strings"

	"github.com/nskforward/playbook/conn"
)

func FileDelete(c *conn.Conn, path ...string) {
	c.Execute(fmt.Sprintf("rm %s", strings.Join(path, " ")))
}

func DirDelete(c *conn.Conn, path string) {
	c.Execute(fmt.Sprintf("rm -rf %s", path))
}
