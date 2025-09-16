package playbook

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"

	"golang.org/x/crypto/ssh"
)

type Conn struct {
	client     *ssh.Client
	addr       string
	user       string
	authMethod []ssh.AuthMethod
	cmd        Command
}

func Connect(opts ...Opt) *Conn {
	conn := &Conn{
		authMethod: make([]ssh.AuthMethod, 0, 4),
	}

	for _, op := range opts {
		op(conn)
	}

	if conn.addr == "" {
		conn.addr = AskStr("addr", false)
	}

	if conn.user == "" {
		conn.user = AskStr("username", false)
	}

	if len(conn.authMethod) == 0 {
		WithPass(AskPass("password"))(conn)
	}

	if !strings.Contains(conn.addr, ":") {
		conn.addr = conn.addr + ":22"
	}

	client, err := ssh.Dial("tcp", conn.addr, &ssh.ClientConfig{
		User:            conn.user,
		Auth:            conn.authMethod,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})

	Catch(err, "ssh connection failed")

	conn.client = client

	fmt.Println("+ connected")

	conn.detectOSRelease()

	return conn
}

func (conn *Conn) Close() {
	conn.client.Close()
}

func (conn *Conn) detectOSRelease() {
	m := getOSRelease(conn.client)
	releases := m["ID_LIKE"]
	if releases == "" {
		Catch(errors.New("cat /etc/os-release | grep ID_LIKE"), "cannot detect OS family")
	}
	release := strings.Split(releases, " ")
	if slices.Contains(release, "debian") {
		fmt.Println("os family: debian")
		conn.cmd = NewDebianCommand(conn.client)
		return
	}
	if slices.Contains(release, "rhel") || slices.Contains(release, "fedora") {
		fmt.Println("os family: rhel")
		conn.cmd = NewRHELCommand(conn.client)
		return
	}

	Catch(errors.New(releases), "unknown os family")
}

func (conn *Conn) Command() Command {
	return conn.cmd
}

type Opt func(conn *Conn)

func WithAddr(addr string) Opt {
	return func(conn *Conn) {
		conn.addr = addr
	}
}

func WithUser(user string) Opt {
	return func(conn *Conn) {
		conn.user = user
	}
}

func WithPass(pass string) Opt {
	return func(conn *Conn) {
		conn.authMethod = append(conn.authMethod, ssh.Password(pass))
	}
}

func WithKey(privateKey string) Opt {
	return func(conn *Conn) {
		data, err := os.ReadFile(privateKey)
		Catch(err, "cannot read the private key file")
		signer, err := ssh.ParsePrivateKey(data)
		Catch(err, "cannot parse private key")
		conn.authMethod = append(conn.authMethod, ssh.PublicKeys(signer))
	}
}
