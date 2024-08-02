package util

import (
	"fmt"
	"strings"

	"golang.org/x/term"
)

func AskString(prompt string) string {
	var answer string
	fmt.Printf("- %s: ", prompt)
	fmt.Scanln(&answer)
	return strings.TrimSpace(answer)
}

func AskStringIfEmpty(prompt, value string) string {
	fmt.Printf("- %s: ", prompt)
	if value != "" {
		fmt.Println(value)
		return value
	}
	var answer string
	fmt.Scanln(&answer)
	return strings.TrimSpace(answer)
}

func AskPassword(prompt string) string {
	fmt.Printf("- %s: ", prompt)
	pass, err := term.ReadPassword(0)
	fmt.Println()
	Check(err)
	return string(pass)
}

func AskPasswordIfEmpty(prompt, value string) string {
	fmt.Printf("- %s: ", prompt)
	if value != "" {
		fmt.Println("****************")
		return value
	}
	pass, err := term.ReadPassword(0)
	fmt.Println()
	Check(err)
	return string(pass)
}
