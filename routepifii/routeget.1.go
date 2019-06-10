package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"net/url"
	"strings"
)

var DEFAULT_IP_URL = "http://192.168.2.1/cgi-bin/luci/admin/wan_ip_get"
var DEFAULT_MAC = "78A351323857"

func httpPostForm() {
	resp, err := http.PostForm(DEFAULT_IP_URL,
		url.Values{"key": {"Value"}, "id": {"123"}})

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}

func httpDo() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", DEFAULT_IP_URL, strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func main() {
	//httpPostForm()
	httpDo()
	// data := "loginId=149832166%40126.com&password=aaa123&exp=-1" //设置提交数据

	// q1, e := http.NewRequest("POST", DEFAULT_IP_URL, strings.NewReader(data))
	// if e != nil {
	// 	fmt.Println("connect not")
	// 	return
	// }
	// q1.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8") //设置Content-Type

	// //	q1.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// h1, e := http.DefaultClient.Do(q1)
	// if e != nil {
	// 	fmt.Println("error")
	// 	return
	// }
	// defer h1.Body.Close()

	// body, err := ioutil.ReadAll(h1.Body)
	// if err != nil {
	// 	// handle error
	// }

	// fmt.Println(string(body))

	// fmt.Println("jiang")
}
