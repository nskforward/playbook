package test

import (
	"fmt"
	"testing"

	"github.com/nskforward/playbook/command"
)

/*
	LINUX:
	echo "some text" | sha1sum
	a5c341bec5c89ed16758435069e3124b3685ad93  -

	openssl sha1 1.txt
	SHA1(1.txt)= a5c341bec5c89ed16758435069e3124b3685ad93

	HEX
	local: 736f6d652074657874
	remote:736f6d6520746578740a
*/

func TestFileHash(t *testing.T) {
	hash, err := command.LocalFileSha1("1.txt")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(hash)
}
