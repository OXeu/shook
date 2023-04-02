package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"math/rand"
)

func GetRandomString(n int) string {
	randBytes := make([]byte, n/2)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}

func getToken() (string, string) {
	db, err := leveldb.OpenFile(DbPath(), nil)
	if err != nil {
		return "", "Database open error\n" + err.Error()
	}
	defer db.Close()
	data, err := db.Get([]byte("token"), nil)
	if err != nil {
		return "", ""
	}
	return string(data), ""
}

func getUrl() string {
	db, err := leveldb.OpenFile(DbPath(), nil)
	if err != nil {
		println("Shook internal error: open database failed\n", err.Error())
		return ""
	}
	defer db.Close()
	url, err := db.Get([]byte("url"), nil)
	if err != nil {
		println("Please run `shook init <url>` to init shook.\n", err.Error())
		return ""
	}
	return string(url)
}
