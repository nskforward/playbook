package util

import (
	"fmt"
	"os"
)

func Check(err error) {
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
}
