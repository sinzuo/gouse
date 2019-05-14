// tcpclient
package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
)

func main() {
	var r1 []byte
	var e error
	t2, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	t1, _ := net.DialTCP("tcp", nil, t2)
	defer t1.Close()
	if len(os.Args) < 2 {
		fmt.Println("please input file")
	} else {
		r1, e = ioutil.ReadFile(os.Args[1])
		if e != nil {
			fmt.Println("file not exiest")
			return
		}
	}
	m1 := make([]byte, 10)

	m2 := strconv.Itoa(len(os.Args[1]))

	copy(m1, m2)
	fmt.Println(m1)
	t1.Write(m1)

	t1.Write([]byte(os.Args[1]))

	t1.Write(r1)

	fmt.Println("Hello World!")
}
