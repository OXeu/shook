package main

import (
	"github.com/syndtr/goleveldb/leveldb"
)

func getToken() (string, string) {
	db, err := leveldb.OpenFile(DbConfigPath(), nil)
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
