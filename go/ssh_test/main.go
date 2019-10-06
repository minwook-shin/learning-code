package main

import (
	"time"

	"github.com/gliderlabs/ssh"
)

func main() {
	MaxTimeout := 20 * time.Second
	IdleTimeout := 10 * time.Second

	ssh.Handle(func(s ssh.Session) {
		i := 0
		for {
			i++
			select {
			case <-time.After(time.Second):
				continue
			case <-s.Context().Done():
				return
			}
		}
	})

	server := &ssh.Server{
		Addr:        ":2222",
		MaxTimeout:  MaxTimeout,
		IdleTimeout: IdleTimeout,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
