package command

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/nskforward/playbook/conn"
	"github.com/nskforward/playbook/util"
)

func FilesEqual(c *conn.Conn, local, remote string) bool {
	if local == "" || remote == "" {
		return false
	}

	if !FileExists(c, remote) {
		return false
	}

	localHash, err := LocalFileSha1(local)
	util.Check(err)
	remoteHash := RemoteFileSha1(c, remote)
	return localHash == remoteHash
}

func RemoteFileSha1(c *conn.Conn, path string) string {
	output := c.Execute(fmt.Sprintf("sha1sum -b %s", path))
	if output == "" {
		util.Check(fmt.Errorf("incorrect sha1sum output: %s", output))
	}
	hash, _, found := strings.Cut(output, " ")
	if !found {
		util.Check(fmt.Errorf("incorrect sha1sum output: %s", output))
	}
	return hash
}

func LocalFileSha1(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := sha1.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
