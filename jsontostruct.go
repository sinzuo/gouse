// structtosql
package main

import (
	//	"bytes"
	//	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type A struct {
	Id         int `key:"1"`
	Name       string
	Password   string
	Department string
	email      string
}

type Btype struct {
	Key    int
	TValue string
}

func jiexi(a interface{}, b map[string]*Btype) (int, int, string, string) {
	a1 := reflect.TypeOf(a)
	var i = 0
	var key = 0
	var value = ""
	for ; i < a1.NumField(); i++ {
		temp := &Btype{}
		temp.TValue = a1.Field(i).Type.Name()
		fmt.Println(a1.Field(i).Tag.Get("key"))
		if a1.Field(i).Tag.Get("key") == "1" {
			fmt.Println("ok")
			temp.Key = 1
			key = 1
			value = a1.Field(i).Name
		}
		b[a1.Field(i).Name] = temp
		fmt.Println(a1.Field(i).Name)
	}

	return i, key, a1.Name(), value
}

var ss = `{"key":"jiang","age":1}`

func main() {
	v1 := make(map[string]interface{})
	s1, _ := os.OpenFile("A.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	defer s1.Close()
	fmt.Fprintln(s1, "type A struct{")
	json.Unmarshal([]byte(ss), &v1)
	for k, v := range v1 {
		m := []byte(k)
		m[0] = m[0] - 32
		fmt.Println(k, reflect.TypeOf(v).Name())
		fmt.Fprintln(s1, "     "+string(m), reflect.TypeOf(v).Name())
	}
	fmt.Fprintln(s1, "}")

	fmt.Println("Hello World!")
}
