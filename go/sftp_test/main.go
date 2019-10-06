package main

import (
	"fmt"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func main() {
	var con *ssh.Client
	client, err := sftp.NewClient(con)
	if err != nil {
		panic(err)
	}

	defer client.Close()

	walk := client.Walk("~/")
	for walk.Step() {
		if walk.Err() != nil {
			continue
		}
		fmt.Println(walk.Path())
	}

	file, err := client.Create("test.txt")
	if err != nil {
		panic(err)
	}

	text := []byte("Hello world")
	if _, err := file.Write(text); err != nil {
		panic(err)
	}

	out, err := sftp.Lstat("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
