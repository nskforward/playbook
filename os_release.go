package playbook

import (
	"bufio"
	"bytes"

	"golang.org/x/crypto/ssh"
)

func getOSRelease(client *ssh.Client) map[string]string {
	cmd := "cat /etc/os-release"
	output, err := execute(client, cmd, false)
	Catch(err, cmd)

	result := make(map[string]string)

	scanner := bufio.NewScanner(bytes.NewReader(output))
	for scanner.Scan() {
		line := scanner.Bytes()
		key, value, ok := bytes.Cut(line, []byte{'='})
		if !ok {
			continue
		}
		if len(value) > 2 && value[0] == '"' && value[len(value)-1] == '"' {
			value = value[1 : len(value)-1]
		}
		result[string(key)] = string(value)
	}

	return result
}
