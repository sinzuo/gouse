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

var ss = `{"persion":[{"mm":"ss"}]}`

//var ss = `{"key":"jiang","age":1,"persion":{"mm":"ss"}}`

func xunlen(s1 *os.File, index int, name string, v1 map[string]interface{}) {
	if index == 1 {
		fmt.Fprintln(s1, "type "+name+" struct {")
	} else {
		fmt.Fprintln(s1, "     "+name+" struct {")
	}
	index++
	for k, v := range v1 {
		m := []byte(k)
		if m[0] >= 97 && m[0] <= 122 {
			m[0] = m[0] - 32
		}
		k = string(m)
		fmt.Println(k, reflect.TypeOf(v).Name())

		t3 := reflect.TypeOf(v)

		v3 := reflect.ValueOf(v)
		fmt.Println(t3.Kind())
		if t3.Kind() == reflect.Map {
			xunlen(s1, index, k, v3.Interface().(map[string]interface{}))

		} else if t3.Kind() == reflect.Slice {

			if v3.Len() >= 1 {
				xunlen(s1, index, k, v3.Index(0).Interface().(map[string]interface{}))
				fmt.Fprint(s1, "[]")
			}
		} else if t3.Kind() == reflect.String {

			fmt.Fprintln(s1, "     "+k, "string")
		} else if t3.Kind() == reflect.Float64 {

			fmt.Fprintln(s1, "     "+k, "int")
		}
	}
	if index == 2 {
		fmt.Fprintln(s1, "}")
	} else {
		fmt.Fprintln(s1, "     }")
	}

}

func main() {
	v1 := make(map[string]interface{})
	s1, _ := os.OpenFile("A.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	defer s1.Close()

	json.Unmarshal([]byte(ss), &v1)
	xunlen(s1, 1, "A", v1)
	fmt.Println("Hello World!")
}
