package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"os/exec"
)

func run(key string) string {
	println("Run hook /" + key)
	db, err := leveldb.OpenFile("hooks", nil)
	if err != nil {
		return "Database open error\n" + err.Error()
	}
	defer db.Close()
	data, err := db.Get([]byte(key), nil)
	if err != nil {
		return "Invoke /" + key + " error\n" + err.Error()
	}
	cmd := exec.Command("sh", "-c", string(data))
	err = cmd.Run()
	if err != nil {
		return "$ " + string(data) + "\nRun /" + key + " command error\n" + err.Error()
	}
	return "Successfully!"
}
func add(key string, pwd string, shell string) string {
	println("Create hook /" + key)
	db, err := leveldb.OpenFile("hooks", nil)
	if err != nil {
		return "Database open error\n" + err.Error()
	}
	defer db.Close()
	var exist = true
	_, e := db.Get([]byte(key), nil)
	if e != nil {
		exist = false
	}

	cmd := "cd " + pwd + " && " + shell
	err = db.Put([]byte(key), []byte(cmd), nil)
	if err != nil {
		return "Create /" + key + " error\n" + err.Error()
	}
	if exist {
		return "/" + key + " hooks updated!\n$ " + cmd
	}
	return "/" + key + " hooks created!\n$ " + cmd
}
func del(key string) string {
	println("Delete hook /" + key)
	db, err := leveldb.OpenFile("hooks", nil)
	if err != nil {
		return "Database open error\n" + err.Error()
	}
	defer db.Close()
	err = db.Delete([]byte(key), nil)
	if err != nil {
		return "Delete /" + key + " error\n" + err.Error()
	}
	return "/" + key + " hooks deleted!\n"
}

func ls() string {
	println("List hooks")
	db, err := leveldb.OpenFile("hooks", nil)
	if err != nil {
		return "Database open error\n" + err.Error()
	}
	ls := fmt.Sprintln("name\tscript")
	defer db.Close()
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		ls += fmt.Sprintf("%s\t%s", string(iter.Key()), string(iter.Value()))
	}
	return ls
}
