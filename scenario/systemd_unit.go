package scenario

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/nskforward/playbook/cmd"
	"github.com/nskforward/playbook/conn"
	"github.com/nskforward/playbook/util"
)

func SystemdUnitCreate(c *conn.Conn, name, exeFilePath, envFilePath, user, group string) {
	fmt.Println("# CREATE SYSTEMD UNIT")

	if exeFilePath == "" {
		util.Check(fmt.Errorf("exe file path cannot be empty"))
	}

	if !strings.HasSuffix(name, ".service") {
		name = name + ".service"
	}

	path := filepath.Join("/etc/systemd/system", name)

	fmt.Println("--> check if unit already exists")

	if cmd.FileExists(c, path) {
		fmt.Println("<-- exists")
		return
	}

	fmt.Println("<-- does not exist")

	fmt.Println("--> create file")
	cmd.FileWrite(c, false, path, getUnit(exeFilePath, envFilePath, user, group))
	fmt.Println("<-- ok")
}

func getUnit(exeFilePath, envFilePath, user, group string) string {
	var buf bytes.Buffer
	buf.WriteString("[Unit]\n")
	buf.WriteString("After=multi-user.target\n\n")
	buf.WriteString("[Service]\n")
	buf.WriteString("ExecStart=")
	buf.WriteString(exeFilePath)
	buf.WriteString("\n")
	buf.WriteString("Type=simple\n")
	buf.WriteString("Restart=always\n")
	buf.WriteString("RestartSec=5\n")
	if envFilePath != "" {
		buf.WriteString("EnvironmentFile=")
		buf.WriteString(envFilePath)
		buf.WriteString("\n")
	}
	if user != "" {
		buf.WriteString("User=")
		buf.WriteString(user)
		buf.WriteString("\n")
	}
	if group != "" {
		buf.WriteString("Group=")
		buf.WriteString(group)
		buf.WriteString("\n")
	}
	buf.WriteString("\n")
	buf.WriteString("[Install]\n")
	buf.WriteString("WantedBy=multi-user.target\n")
	return buf.String()
}
