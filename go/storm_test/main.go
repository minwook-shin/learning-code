package main

import (
	"fmt"
	"time"

	"github.com/asdine/storm"
)

type testInfo struct {
	ID        int       `storm:"id,increment"`
	Email     string    `storm:"unique"`
	CreatedAt time.Time `storm:"index"`
}

func main() {
	db, err := storm.Open("test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Drop(&testInfo{})
	if err != nil {
		fmt.Println(err)
	}

	err = db.Init(&testInfo{})
	if err != nil {
		fmt.Println(err)
	}

	test1 := testInfo{
		Email:     "test@example.com",
		CreatedAt: time.Now(),
	}

	err = db.Save(&test1)
	if err != nil {
		fmt.Println(err)
	}

	var test2 testInfo
	err = db.One("Email", "test@example.com", &test2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(test2)

	err = db.Update(&testInfo{ID: 1, Email: "new@example.com"})
	if err != nil {
		fmt.Println(err)
	}

	var testList []testInfo
	err = db.All(&testList)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(testList)

	err = db.DeleteStruct(&test1)
	if err != nil {
		fmt.Println(err)
	}

	err = db.All(&testList)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(testList)
}
