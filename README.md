# playbook
Automation of remote server commands over SSH

## example

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

Then the make file can have very short rules

```
...
deploy:
    go run cmd/script/deploy.go
...
```

You can explore the folders 'scenario' and 'cmd' to find the supported commands.

Feel free to create an issue if something went wrong or you have a suggestion.