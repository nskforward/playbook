package playbook

import (
	"fmt"
	"strings"

	"golang.org/x/term"
)

func AskStr(prompt string, allowEmpty bool) string {
	var answer string
	for {
		fmt.Printf("- %s: ", prompt)
		fmt.Scanln(&answer)
		answer = strings.TrimSpace(answer)
		if allowEmpty || answer != "" {
			break
		}
	}
	return answer
}

func AskPass(prompt string) string {
	fmt.Printf("- %s: ", prompt)
	pass, err := term.ReadPassword(0)
	fmt.Println()
	Catch(err, "invalid password")
	return string(pass)
}
