package playbook

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func GetLocalFile(path string) []byte {
	output, err := os.ReadFile(path)
	Catch(err, fmt.Sprintf("cannot read file: %s", path))
	return output
}

func GeneratePassword(prefix string, length int) string {
	allowed := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!?_-*"
	generated := make([]byte, length)
	for i := range generated {
		b := allowed[rand.Intn(len(allowed))]
		generated[i] = b
	}
	return strings.Join([]string{prefix, string(generated)}, "")
}
