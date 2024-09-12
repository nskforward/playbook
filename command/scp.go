package command

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/nskforward/playbook/conn"
	"github.com/nskforward/playbook/util"
)

func ScpFile(c *conn.Conn, localFilePath, remoteFilePath string) bool {
	if FilesEqual(c, localFilePath, remoteFilePath) {
		return false
	}

	fmt.Println("uploading...", localFilePath, "-->", remoteFilePath)

	sftp := c.SFTP()
	defer sftp.Close()

	src, err := os.Open(localFilePath)
	util.Check(err)
	defer src.Close()

	dst, err := sftp.Create(remoteFilePath)
	util.Check(err)
	defer dst.Close()

	_, err = io.Copy(dst, src)
	util.Check(err)

	return true
}

func ScpDir(c *conn.Conn, localDirPath, remoteDirPath, owner string) {

	items, err := os.ReadDir(localDirPath)
	util.Check(err)

	if !DirExists(c, remoteDirPath) {
		DirCreate(c, false, remoteDirPath)
		Chown(c, true, owner, owner, remoteDirPath)
		Chmod(c, true, remoteDirPath, util.NewPerm(7, 7, 0))
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
