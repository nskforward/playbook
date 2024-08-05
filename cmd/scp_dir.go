package cmd

import (
	"os"
	"path/filepath"

	"github.com/nskforward/playbook/conn"
	"github.com/nskforward/playbook/util"
)

func ScpDir(c *conn.Conn, localDirPath, remoteDirPath, owner string) {

	items, err := os.ReadDir(localDirPath)
	util.Check(err)

	if !DirExists(c, remoteDirPath) {
		DirMake(c, false, remoteDirPath)
		Chown(c, true, owner, owner, remoteDirPath)
	}

	for _, item := range items {

		src := filepath.Join(localDirPath, item.Name())
		dst := filepath.Join(remoteDirPath, item.Name())

		if item.IsDir() {
			ScpDir(c, src, dst, owner)
			continue
		}
		if item.Type().IsRegular() {
			ScpFile(c, src, dst)
		}
	}
}
