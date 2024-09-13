[![Go Reference](https://pkg.go.dev/badge/github.com/nskforward/playbook?status.svg)](https://pkg.go.dev/github.com/nskforward/playbook?tab=doc)

# Playbook

Automation of remote server commands over SSH

> [!WARNING]
> Use in production with caution and only after proper testing in a development/test environment.

## Installation
```
go get -u github.com/nskforward/playbook
```

## Example

You can create a complex deployment script using the following Golang project structure

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

You can explore the folders 'command' to find the supported commands.

Feel free to create an issue if something went wrong or you have a suggestion.
