package conn

import (
	"fmt"
	"strings"

	"github.com/nskforward/playbook/util"
)

type Config struct {
	Addr    string
	User    string
	Key     string
	Pass    string
	HasSudo bool
	Debug   bool
}

type LoginArg func(cfg *Config)

func Key(path string) LoginArg {
	return func(cfg *Config) {
		cfg.Key = path
	}
}

func Sudo(cfg *Config) {
	cfg.HasSudo = true
}

func Debug(cfg *Config) {
	cfg.Debug = true
}

func Connect(addr, user string, args ...LoginArg) *Conn {
	cfg := Config{
		Addr: addr,
		User: user,
	}
	for _, f := range args {
		f(&cfg)
	}

	fmt.Println("# LOGIN")

	cfg.Addr = util.AskStringIfEmpty("host/ip[:port]", cfg.Addr)
	if cfg.Addr == "" {
		util.Check(fmt.Errorf("host must be specified"))
	}

	cfg.User = util.AskStringIfEmpty("user", cfg.User)
	if cfg.User == "" {
		util.Check(fmt.Errorf("user must be specified"))
	}
	if cfg.Key == "" {
		cfg.Pass = util.AskPassword("password")
	}

	if !strings.Contains(cfg.Addr, ":") {
		cfg.Addr = cfg.Addr + ":22"
	}

	fmt.Printf("try to ssh connect to %s@%s\n", cfg.User, cfg.Addr)
	util.Confitm()

	fmt.Println("*")

	conn := dial(cfg)

	fmt.Println("--------------------------")
	fmt.Println("| successfully connected |")
	fmt.Println("--------------------------")

	fmt.Println("*")

	return conn
}
