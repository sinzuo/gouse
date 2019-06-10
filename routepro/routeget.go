package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var DEFAULT_IP_URL = "http://192.168.3.242:8081/wifi/acsjson"
var DEFAULT_MAC = "78A351323857"

func main() {
	q1, e := http.NewRequest("POST", DEFAULT_IP_URL, strings.NewReader("name=cjb"))
	if e != nil {
		fmt.Println("connect not")
		return
	}
	//	q1.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	q1.Header.Set("Cookie", "name=anny")
	u1 := url.Values{"url": {"wo", "fang"}}

	q1.Form = u1

	h1, e := http.DefaultClient.Do(q1)
	if e != nil {
		fmt.Println("error")
		return
	}
	defer h1.Body.Close()

	body, err := ioutil.ReadAll(h1.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

	fmt.Println("jiang")
}
