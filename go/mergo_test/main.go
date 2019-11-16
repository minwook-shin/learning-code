package main

import (
	"fmt"

	"github.com/imdario/mergo"
)

type User struct {
	ID       string
	Password string
}

func main() {
	newInfo := User{
		ID:       "user_id_1",
		Password: "0000",
	}
	oldInfo := User{
		ID: "user_id_2",
	}
	mergo.Merge(&oldInfo, newInfo)
	fmt.Println(oldInfo)

	mergo.Merge(&oldInfo, newInfo, mergo.WithOverride)
	fmt.Println(oldInfo)

	mergo.Merge(&oldInfo, newInfo, mergo.WithOverride, mergo.WithTypeCheck)
	fmt.Println(oldInfo)
}
