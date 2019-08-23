package main

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

func decode(b []byte) string {
	return string(b[:len(b)])
}

func main() {
	db, err := leveldb.OpenFile("test_path/test_file", nil)
	if err != nil {
		panic(err)
	}
	err = db.Put([]byte("key"), []byte("value"), nil)
	data, err := db.Get([]byte("key"), nil)
	fmt.Println(decode(data))

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Println(key, decode(value))
	}
	iter.Release()
	err = iter.Error()

	err = db.Delete([]byte("key"), nil)

	batch := new(leveldb.Batch)
	batch.Put([]byte("key1"), []byte("value1"))
	batch.Put([]byte("key2"), []byte("value2"))
	batch.Delete([]byte("key1"))
	err = db.Write(batch, nil)
	defer db.Close()

	o := &opt.Options{
		Filter: filter.NewBloomFilter(10),
	}
	db, err = leveldb.OpenFile("test_path/test_BloomFilter", o)
	if err != nil {
		panic(err)
	}

	defer db.Close()
}
