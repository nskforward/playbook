package conn

import (
	"fmt"
	"strings"

	"github.com/nskforward/playbook/util"
)

type Config struct {
	Host    string
	User    string
	Key     string
	Pass    string
	HasSudo bool
	Port    int
}

type LoginArg func(cfg *Config)

func Key(path string) LoginArg {
	return func(cfg *Config) {
		cfg.Key = path
	}
}

func Port(port int) LoginArg {
	return func(cfg *Config) {
		cfg.Port = port
	}
}

func Sudo() LoginArg {
	return func(cfg *Config) {
		cfg.HasSudo = true
	}
}

func Connect(host, user string, args ...LoginArg) *Conn {
	cfg := Config{
		Host: host,
		User: user,
		Port: 22,
	}
	for _, f := range args {
		f(&cfg)
	}

	fmt.Println("# LOGIN")

	cfg.Host = util.AskStringIfEmpty("host(ip)", cfg.Host)
	if cfg.Host == "" {
		util.Check(fmt.Errorf("host must be specified"))
	}
	if !strings.Contains(cfg.Host, ":") {
		cfg.Host += ":22"
	}
	cfg.User = util.AskStringIfEmpty("user", cfg.User)
	if cfg.User == "" {
		util.Check(fmt.Errorf("user must be specified"))
	}
	if cfg.Key == "" {
		cfg.Pass = util.AskPassword("password")
	}

	fmt.Printf("try to ssh connect to %s@%s\n", cfg.User, cfg.Host)
	util.Confitm()

	conn := dial(cfg)

	fmt.Println("successfully connected")
	return conn
}
