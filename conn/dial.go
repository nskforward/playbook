package conn

import (
	"fmt"
	"os"

	"github.com/nskforward/playbook/util"
	"golang.org/x/crypto/ssh"
)

func dial(cfg Config) *Conn {
	client, err := ssh.Dial("tcp", cfg.Host, &ssh.ClientConfig{
		User:            cfg.User,
		Auth:            []ssh.AuthMethod{detectAuthMethod(cfg.Key, cfg.Pass)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	util.Check(err)
	os := GetOS(client)
	fmt.Println("detected os family:", os)
	return &Conn{client: client, os: os, sudo: cfg.HasSudo}
}

func detectAuthMethod(keyPath, pass string) ssh.AuthMethod {
	if keyPath == "" {
		return ssh.Password(pass)
	}
	pemBytes, err := os.ReadFile(keyPath)
	util.Check(err)
	signer, err := ssh.ParsePrivateKey(pemBytes)
	util.Check(err)
	return ssh.PublicKeys(signer)
}
