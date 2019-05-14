// tcpclient
package main

import (
	"fmt"
	"net"
)

func main() {
	t2, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	t1, _ := net.DialTCP("tcp", nil, t2)

	t1.Write([]byte("ok"))
	fmt.Println("Hello World!")
}
