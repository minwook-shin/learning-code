package main

import (
	"fmt"

	"github.com/tidwall/buntdb"
)

func main() {
	db, err := buntdb.Open("data.db")
	// db, err := buntdb.Open(":memory:")
	if err != nil {
		panic(err)
	}

	err = db.Update(func(tx *buntdb.Tx) error {
		_, _, err = tx.Set("first_key", "hello", nil)
		_, _, err = tx.Set("second_key", "world", nil)
		return nil
	})

	err = db.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get("first_key")
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", val)
		return nil
	})

	err = db.View(func(tx *buntdb.Tx) error {
		err = tx.Ascend("", func(key, value string) bool {
			fmt.Printf("%s, %s\n", key, value)
			return true
		})
		return err
	})

	db.CreateIndex("tests", "test:*", buntdb.IndexString)
	err = db.Update(func(tx *buntdb.Tx) error {
		_, _, err = tx.Set("test:0:tests", "fisrt_test", nil)
		_, _, err = tx.Set("production:1:tests", "real", nil)
		_, _, err = tx.Set("test:2:tests", "second_test", nil)
		return nil
	})

	db.View(func(tx *buntdb.Tx) error {
		tx.Ascend("tests", func(key, val string) bool {
			fmt.Printf("%s %s\n", key, val)
			return true
		})
		return nil
	})
	defer db.Close()
}
