package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/jlaffaye/ftp"
)

func main() {
	connect, err := ftp.Dial("127.0.0.1:21", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		panic(err)
	}

	err = connect.Login("USER", "PASSWORD")
	if err != nil {
		panic(err)
	}

	buffer := bytes.NewBufferString("Hello")
	err = connect.Stor("test.txt", buffer)
	if err != nil {
		panic(err)
	}

	response, err := connect.Retr("test.txt")
	if err != nil {
		panic(err)
	}

	byteData, err := ioutil.ReadAll(response)
	fmt.Println(string(byteData))

	if err := connect.Quit(); err != nil {
		panic(err)
	}
}
