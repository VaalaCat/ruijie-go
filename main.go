package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func getOnlineInfo() bool {
	_, err := http.Get("https://connect.rom.miui.com/generate_204")
	if err != nil {
		fmt.Printf("Start Login Process\n")
		return false
	}
	fmt.Println("Already Online")
	return true
}

func getQueryString() string {
	fmt.Println("Trying to get query string")
	resp, _ := http.Get("http://baidu.com")
	body, _ := ioutil.ReadAll(resp.Body)

	bodyStr := string(body)
	cut := strings.Index(bodyStr, "wlanuser")
	bodyStr = bodyStr[cut:]
	cut = strings.Index(bodyStr, "'")
	bodyStr = bodyStr[:cut]
	queryString := bodyStr
	return queryString
}

func login(queryString string, userId string, password string) {
	urlValues := url.Values{
		"queryString": {queryString},
		"userId":      {userId},
		"password":    {password},
	}
	reqBody := urlValues.Encode()
	resp, _ := http.Post("http://172.26.156.158/eportal/InterFace.do?method=login", "application/x-www-form-urlencoded", strings.NewReader(reqBody))
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Printf("Usage ./ruijie-go 200000 20000\n")
		os.Exit(1)
	}
	if !getOnlineInfo() {
		queryString := getQueryString()
		login(queryString, args[1], args[2])
	}
}
