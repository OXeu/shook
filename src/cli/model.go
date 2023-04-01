package main

import (
	"bytes"
	"github.com/mitchellh/go-homedir"
	"github.com/syndtr/goleveldb/leveldb"
	"io"
	"net/http"
	url2 "net/url"
	"path/filepath"
)

const Path = "~/.config/shook/config"

func cmdPwd() {
	abs, err := filepath.Abs(DbPath())
	if err != nil {
		println("Open file error\n" + err.Error())
		return
	}
	println(abs)
}
func DbPath() string {
	expand, err := homedir.Expand(Path)
	if err != nil {
		panic("get home dir error\n" + err.Error())
		return ""
	}
	return expand
}

func cmdInit(url string) {
	if len(url) == 0 {
		println("Base URL can not be empty")
		return
	}
	db, err := leveldb.OpenFile(DbPath(), nil)
	if err != nil {
		println("Shook init error: open database failed\n", err.Error())
		return
	}
	err = db.Put([]byte("url"), []byte(url), nil)
	if err != nil {
		println("Shook init error: set url failed\n", err.Error())
		return
	}
	println("Shook init succeeded!")
}

func getUrl() string {
	db, err := leveldb.OpenFile(DbPath(), nil)
	if err != nil {
		println("Shook internal error: open database failed\n", err.Error())
		return ""
	}
	url, err := db.Get([]byte("url"), nil)
	if err != nil {
		println("Please run `shook init <url>` to init shook.\n", err.Error())
		return ""
	}
	return string(url)
}

func cmdCreate(key string, pwd string, shell string) {
	if len(key) == 0 {
		println("name can not be empty")
		return
	}
	if len(pwd) == 0 {
		println("work dir is empty")
		return
	}
	if len(shell) == 0 {
		println("shell can not be empty")
		return
	}
	url := getUrl()
	if len(url) == 0 {
		return
	}
	client := http.Client{}
	formValues := url2.Values{}
	formValues.Set("pwd", pwd)
	formValues.Set("shell", shell)
	formDataStr := formValues.Encode()
	formDataBytes := []byte(formDataStr)
	formBytesReader := bytes.NewReader(formDataBytes)
	response, err := client.Post(url+"/admin/"+key, "application/x-www-form-urlencoded", formBytesReader)
	if err != nil {
		println("Request failed\n", err.Error())
		return
	}
	all, err := io.ReadAll(response.Body)
	if err != nil {
		println("Response read failed\n", err.Error())
		return
	}
	println(string(all))
}
func cmdDel(key string) {
	url := getUrl()
	if len(url) == 0 {
		return
	}
	req, _ := http.NewRequest("DELETE", url+"/admin/"+key, nil)
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		println("Request failed\n", err.Error())
		return
	}
	all, err := io.ReadAll(response.Body)
	if err != nil {
		println("Response read failed\n", err.Error())
		return
	}
	println(string(all))
}

func cmdRun(key string) {
	url := getUrl()
	if len(url) == 0 {
		return
	}
	client := http.Client{}
	response, err := client.Get(url + "/" + key)
	if err != nil {
		println("Request failed\n", err.Error())
		return
	}
	all, err := io.ReadAll(response.Body)
	if err != nil {
		println("Response read failed\n", err.Error())
		return
	}
	println(string(all))
}

func cmdLs() {
	url := getUrl()
	if len(url) == 0 {
		return
	}
	client := http.Client{}
	response, err := client.Get(url + "/admin")
	if err != nil {
		println("Request failed\n", err.Error())
		return
	}
	all, err := io.ReadAll(response.Body)
	if err != nil {
		println("Response read failed\n", err.Error())
		return
	}
	println(string(all))
}
