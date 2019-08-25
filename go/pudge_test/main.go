package main

import (
	"fmt"

	"github.com/recoilme/pudge"
)

func main() {
	pudge.Set("test_path/test_file", "key", "value")

	variable := ""
	pudge.Get("test_path/test_file", "key", &variable)
	fmt.Println(variable)

	pudge.DeleteFile("test_path/test_file")

	cfg := &pudge.Config{}
	db, err := pudge.Open("test_path/test_file", cfg)
	if err != nil {
		panic(err)
	}

	type TestStruct struct {
		Fisrt  int
		Second int
	}

	db.Set(1, &TestStruct{Fisrt: 1, Second: 1})
	db.Set(2, &TestStruct{Fisrt: 2, Second: 2})

	var test TestStruct
	db.Get(1, &test)
	fmt.Println(test)

	keys, _ := db.Keys(0, 2, 0, true)
	for _, key := range keys {
		var t TestStruct
		db.Get(key, &t)
		fmt.Println(t)
	}

	defer db.DeleteFile()

	if err := pudge.CloseAll(); err != nil {
		panic(err)
	}

}
