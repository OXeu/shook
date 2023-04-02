package main

import (
	"bytes"
	"github.com/mitchellh/go-homedir"
	"github.com/syndtr/goleveldb/leveldb"
	"io"
	"net/http"
	url2 "net/url"
	"path/filepath"
	"strings"
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

func cmdInit(url string, token string) {
	if len(url) == 0 {
		println("Base URL can not be empty")
		return
	}

	if len(token) == 0 {
		token, _ = getToken()
		if len(token) == 0 {
			println("No token was set, generating token")
			token = GetRandomString(32)
		} else {
			println("No token was set, use previous token")
		}
		println(token)
	}

	// verify
	formValues := url2.Values{}
	formValues.Set("token", token)
	formDataStr := formValues.Encode()
	formDataBytes := []byte(formDataStr)
	formBytesReader := bytes.NewReader(formDataBytes)
	req, _ := http.NewRequest("PUT", url+"/admin", formBytesReader)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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
	if strings.Contains(string(all), "successfully") {
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
		err = db.Put([]byte("token"), []byte(token), nil)
		if err != nil {
			println("Shook init error: set token failed\n", err.Error())
			return
		}
		println("Shook initialization successfully!")
	}
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
	req, _ := http.NewRequest("POST", url+"/admin/"+key, formBytesReader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	token, errMsg := getToken()
	if len(errMsg) != 0 {
		println(errMsg)
		return
	}
	req.Header.Set("Authorization", "Bearer "+token)
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
func cmdDel(key string) {
	url := getUrl()
	if len(url) == 0 {
		return
	}
	req, _ := http.NewRequest("DELETE", url+"/admin/"+key, nil)
	token, errMsg := getToken()
	if len(errMsg) != 0 {
		println(errMsg)
		return
	}
	req.Header.Set("Authorization", "Bearer "+token)
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
	req, _ := http.NewRequest("GET", url+"/"+key, nil)
	token, errMsg := getToken()
	if len(errMsg) != 0 {
		println(errMsg)
		return
	}
	req.Header.Set("Authorization", "Bearer "+token)
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

func cmdLs() {
	url := getUrl()
	if len(url) == 0 {
		return
	}
	client := http.Client{}
	req, _ := http.NewRequest("GET", url+"/admin", nil)
	token, errMsg := getToken()
	if len(errMsg) != 0 {
		println(errMsg)
		return
	}
	req.Header.Set("Authorization", "Bearer "+token)
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
