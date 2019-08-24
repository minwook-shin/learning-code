package main

import (
	"fmt"

	"github.com/xujiajun/nutsdb"
)

func main() {
	opt := nutsdb.DefaultOptions
	opt.Dir = "test_path/nutsdb_file"
	db, err := nutsdb.Open(opt)
	if err != nil {
		panic(err)
	}

	err = db.Update(
		func(tx *nutsdb.Tx) error {
			return nil
		})

	if err := db.Update(
		func(tx *nutsdb.Tx) error {
			if err := tx.Put("test_bucket", []byte("test_key"), []byte("test_val"), 0); err != nil {
				return err
			}
			return nil
		}); err != nil {
		panic(err)
	}

	if err := db.View(
		func(tx *nutsdb.Tx) error {
			if e, err := tx.Get("test_bucket", []byte("test_key")); err != nil {
				return err
			} else {
				fmt.Println(string(e.Value))
			}
			return nil
		}); err != nil {
		panic(err)
	}
	if err := db.Update(
		func(tx *nutsdb.Tx) error {
			if err := tx.Delete("test_bucket", []byte("test_key")); err != nil {
				return err
			}
			return nil
		}); err != nil {
		panic(err)
	}

	if err := db.View(
		func(tx *nutsdb.Tx) error {
			if e, err := tx.Get("test_bucket", []byte("test_key")); err != nil {
				return err
			} else {
				fmt.Println(string(e.Value))
			}
			return nil
		}); err != nil {
		fmt.Println(err)
	}

	if err := db.Update(
		func(tx *nutsdb.Tx) error {
			if err := tx.Put("test_bucket", []byte("test_key"), []byte("test_val"), 60); err != nil {
				return err
			}
			return nil
		}); err != nil {
		panic(err)
	}
	if err := db.View(
		func(tx *nutsdb.Tx) error {
			entries, err := tx.GetAll("test_bucket")
			if err != nil {
				return err
			}
			for _, entry := range entries {
				fmt.Println(string(entry.Key), string(entry.Value))
			}
			return nil
		}); err != nil {
		panic(err)
	}

	err = db.Backup("backup")
	if err != nil {
		panic(err)
	}

	defer db.Close()
}
