package main

import (
	"crypto/rand"
	"fmt"
	"os"

	"golang.org/x/net/context"

	libsass "github.com/wellington/go-libsass"
)

func marshal(ctx context.Context, usv libsass.SassValue) (*libsass.SassValue, error) {
	b := make([]byte, 10)
	rand.Read(b)
	res, err := libsass.Marshal(fmt.Sprintf("'%x'", b))
	return &res, err
}

func main() {
	r, err := os.Open("main.scss")
	libsass.RegisterSassFunc("crypto()", marshal)

	libsass.RegisterHeader(`
/*
 * RegisterHeader
 */`)

	if err != nil {
		panic(err)
	}

	compiler, err := libsass.New(os.Stdout, r)

	if err != nil {
		panic(err)
	}

	path := []string{"config"}
	err = compiler.Option(libsass.IncludePaths(path))
	if err != nil {
		panic(err)
	}

	if err := compiler.Run(); err != nil {
		panic(err)
	}

}
