package playbook

import (
	"fmt"
	"os"
)

func Catch(err error, text string) {
	if err != nil {
		fmt.Printf("[ERROR] %s\n\t- %s\n", text, err)
		os.Exit(1)
	}
}
