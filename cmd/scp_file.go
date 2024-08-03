package cmd

import (
	"io"
	"os"

	"github.com/nskforward/playbook/conn"
	"github.com/nskforward/playbook/util"
)

func ScpFile(c *conn.Conn, localFilePath, remoteFilePath string) {
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
}
