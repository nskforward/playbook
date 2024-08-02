package util

import (
	"fmt"
	"os"
)

func Confitm() {
	var answer string
	fmt.Print("- continue? (y/n): ")
	fmt.Scanln(&answer)
	if answer != "y" {
		fmt.Println("finish")
		os.Exit(0)
	}
}
