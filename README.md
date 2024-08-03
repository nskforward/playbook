# playbook
Automation of remote server commands over SSH

> [!WARNING]
> Use in production with caution and only after proper testing in a development/test environment.

## Installation
```
go get -u github.com/nskforward/playbook
```

## Example

You can create a complex deploy script on Go using the following structure as suggestion

```
/root
    /cmd
        /app
            main.go
        /script
            deploy.go
    ...
    makefile
```

makefile

```
...
deploy:
    go run cmd/script/deploy.go
...
```

deploy.go

```go
import (
	"github.com/nskforward/playbook/cmd"
	"github.com/nskforward/playbook/conn"
	"github.com/nskforward/playbook/scenario"
)

func main() {
    addr := "a ssh server host[:port]"
    user := "a ssh user"

    // estabilish connection
    // errors are checked under the hood
    // if an error then the proccess exit with code 1
	c := conn.Connect(addr, user, conn.Debug, conn.Sudo, conn.Key(pathToLocalPrivateKey))
	
    // create a system user 'app' who can't ssh login and does not have a HOME dir
    // if user already exists then just follow to the next step without error
	scenario.UserCreate(c, "app", true)

    // prepare app dir
    if !cmd.DirExists(c, "/app"){
        cmd.DirMake(c, true, "/app")
    }
	
    // upload your app file to remote server
	cmd.ScpFile(c, "./bin/app_linux_amd64", "/app/app_linux_amd64")
	cmd.Chown(c, true, "app", "app", "/app")
	cmd.Chmod(c, true, "/app", util.NewPerm(7, 7, 0))

    // make systemd unit file that pointed to our app
    // if unit already exists then just follow to the next step without error
	scenario.SystemdUnitCreate(c, "api.service", "/app/app_linux_amd64", "", "app", "app")

    // start our service
	cmd.Systemctl(c, "api.service", cmd.Restart)
}
```

The "cmd" folder contains the simple low level commands that do exactly one thing

The "scenario" folder contains the complex scenarios that consist of several low level commands

You can explore the folders 'scenario' and 'cmd' to find the supported commands.

Feel free to create an issue if something went wrong or you have a suggestion.