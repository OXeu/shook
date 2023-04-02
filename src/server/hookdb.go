package main

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/syndtr/goleveldb/leveldb"
	"os/exec"
)

const Path = "~/.config/shook/hooks"
const ConfigPath = "~/.config/shook/server-config"

func DbPath() string {
	expand, err := homedir.Expand(Path)
	if err != nil {
		panic("get home dir error\n" + err.Error())
		return ""
	}
	return expand
}
func DbConfigPath() string {
	expand, err := homedir.Expand(ConfigPath)
	if err != nil {
		panic("get home dir error\n" + err.Error())
		return ""
	}
	return expand
}

func run(key string) string {
	println("Run hook /" + key)
	db, err := leveldb.OpenFile(DbPath(), nil)
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
func initServer(token string) string {
	println("Init hook server with token")
	println(token)
	db, err := leveldb.OpenFile(DbConfigPath(), nil)
	if err != nil {
		return "Database open error\n" + err.Error()
	}
	defer db.Close()
	var exist = true
	t, e := db.Get([]byte("token"), nil)
	if e != nil || len(t) == 0 {
		exist = false
	}
	if err != nil {
		return "Init server error\n" + err.Error()
	}
	if exist {
		return "Invalid Operation. This server has been initialized yet."
	} else {
		err = db.Put([]byte("token"), []byte(token), nil)
		return "Server initialization successfully. Please keep your token carefully."
	}
}
func add(key string, pwd string, shell string) string {
	println("Create hook /" + key)
	db, err := leveldb.OpenFile(DbPath(), nil)
	if err != nil {
		return "Database open error\n" + err.Error()
	}
	defer db.Close()
	var exist = true
	_, e := db.Get([]byte(key), nil)
	if e != nil {
		exist = false
	}

	cmd := "cd " + pwd + " ; " + shell
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
	db, err := leveldb.OpenFile(DbPath(), nil)
	if err != nil {
		return "Database open error\n" + err.Error()
	}
	defer db.Close()
	var exist = true
	_, e := db.Get([]byte(key), nil)
	if e != nil {
		exist = false
	}
	if exist {
		err = db.Delete([]byte(key), nil)
		if err != nil {
			return "Delete /" + key + " error\n" + err.Error()
		}
		return "/" + key + " hooks deleted!"
	} else {
		return "/" + key + " hooks did not exist!"
	}
}

func ls() string {
	println("List hooks")
	db, err := leveldb.OpenFile(DbPath(), nil)
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
