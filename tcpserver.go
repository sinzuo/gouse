// tcpserver
package main

import (
	"fmt"
	"net"
)

func recv(t3 net.Conn) {

	for {
		b1 := make([]byte, 1024)
		_, e := t3.Read(b1)
		if e != nil {
			return
		}
		fmt.Println(b1)

	}

}

//处理连接
func handleConnection(conn net.Conn) {
	buffer := make([]byte, 2048) //建立一个slice
	for {
		n, err := conn.Read(buffer) //读取客户端传来的内容
		if err != nil {
			//           Log(conn.RemoteAddr().String(), "connection error: ", err)
			return //当远程客户端连接发生错误（断开）后，终止此协程。
		}

		fmt.Println(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))

		//返回给客户端的信息
		strTemp := "CofoxServer got msg "
		conn.Write([]byte(strTemp))
	}
}

func main() {

	t1, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	t2, e := net.ListenTCP("tcp", t1)
	if e != nil {
		fmt.Println("ok")
	}
	for {

		t3, _ := t2.Accept()
		go recv(t3)
	}
	fmt.Println("Hello World!")
}
