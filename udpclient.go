// udpclient
package main

import (
	"fmt"
	"net"
)

func main() {
	n1, e := net.Dial("udp", "192.168.3.242:8888")

	if e != nil {
		fmt.Println("error")
		return
	}
	defer n1.Close()
	n1.Write([]byte("hello"))
	n1.Write([]byte("world"))

	fmt.Println("Hello World!")
}
