// udpserver
package main

import (
	"fmt"
	"net"
)

func main() {
	b1 := make([]byte, 2048)
	n1, _ := net.ResolveUDPAddr("udp", ":8888")
	n2, _ := net.ListenUDP("udp", n1)
	for {
		n2.Read(b1)
		fmt.Println(string(b1))
	}
	fmt.Println("Hello World!")
}
