// structtosql
package main

import (
	//	"bytes"
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

/*
CREATE TABLE `userinfo` (
    `uid` INT(10) NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(64)  DEFAULT NULL,
    `password` VARCHAR(32)  DEFAULT NULL,
    `department` VARCHAR(64)  DEFAULT NULL,
    `email` varchar(64) DEFAULT NULL,
    PRIMARY KEY (`uid`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8
*/

var t1 = "CREATE TABLE "
var t2 = " VARCHAR(64)  DEFAULT NULL"
var t3 = ")ENGINE=InnoDB DEFAULT CHARSET=utf8"
var t4 = "INT(10) NOT NULL AUTO_INCREMENT"
var t5 = "    PRIMARY KEY (`"

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

func main() {
	var i = 0

	b1 := make(map[string]*Btype, 20)
	s1, _ := os.OpenFile("A.txt", os.O_CREATE|os.O_RDWR|os.O_SYNC, os.ModePerm)
	defer s1.Close()
	index, key, name, value := jiexi(A{}, b1)

	fmt.Fprintln(s1, t1+" `"+name+"`(")
	for k, v := range b1 {
		i++
		if v.TValue == "string" {

			if i < index || key == 1 {
				fmt.Fprintln(s1, "    `"+k+"` ", t2, ",")
			} else {
				fmt.Fprintln(s1, "    `"+k+"` ", t2)
			}
		} else if v.TValue == "int" {
			if i < index || key == 1 {
				fmt.Fprintln(s1, "    `"+k+"` ", t4, ",")
			} else {
				fmt.Fprintln(s1, "    `"+k+"` ", t4)
			}
		}
	}
	if key == 1 {
		fmt.Fprintln(s1, t5+value+"`)")
	}
	fmt.Fprintln(s1, t3)

	fmt.Println("Hello World!")
}
