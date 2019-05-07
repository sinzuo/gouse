package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {

	//resp, _ := doGet("https://www.baidu.com")
	resp, _ := doPost("http://www.baidu.com", "application/json;charset=utf-8")
	//resp, _ := doPostForm("http://www.baidu.com")
	defer resp.Body.Close() //go的特殊语法，main函数执行结束前会执行resp.Body.Close()

	fmt.Println(resp.StatusCode)          //有http的响应码输出
	if resp.StatusCode == http.StatusOK { //如果响应码为200
		body, err := ioutil.ReadAll(resp.Body) //把响应的body读出
		if err != nil {                        //如果有异常
			fmt.Println(err) //把异常打印
			log.Fatal(err)   //日志
		}
		fmt.Println(string(body)) //把响应的文本输出到console
	}

}

/**
以GET的方式请求
**/
func doGet(url string) (r *http.Response, e error) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(resp.StatusCode)
		fmt.Println(err)
		log.Fatal(err)
	}
	return resp, err
}

/**
以POST的方式请求
**/
type Mafdsadf struct {
}

func (I *Mafdsadf) Read(p []byte) (n int, err error) {
	p = []byte("this is ok")
	return 10, nil

}

/*
Http Header里的Content-Type一般有这三种：
application/x-www-form-urlencoded：数据被编码为名称/值对。这是标准的编码格式。
multipart/form-data： 数据被编码为一条消息，页上的每个控件对应消息中的一个部分。
text/plain
*/

func doPost(url string, bodyType string) (r *http.Response, e error) {

	resp, err := http.Post(url, bodyType, &Mafdsadf{})

	if err != nil {
		fmt.Println(resp.StatusCode)
		fmt.Println(err)
		log.Fatal(err)
	}

	return resp, err
}

/**
以Post表单的方式请求
**/
func doPostForm(urlStr string) (r *http.Response, e error) {

	v := url.Values{"method": {"get"}, "id": {"1"}}
	v.Add("name1", "1")
	v.Add("name2", "2")

	resp, err := http.PostForm(urlStr, v)

	if err != nil {
		fmt.Println(resp.StatusCode)
		fmt.Println(err)
		log.Fatal(err)
	}

	return resp, err

}
