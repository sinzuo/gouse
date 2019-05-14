// tcpserver
package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"
)

func recv(t3 net.Conn) {
	var s1 *os.File
	b1 := make([]byte, 2048)
	n, e := t3.Read(b1)
	defer func() {
		t3.Close()
		s1.Close()
	}()

	if e != nil {
		return
	}
	if n < 10 {
		fmt.Println("package is len< 10")
		return
	}
	fmt.Println(n)

	index := bytes.IndexByte(b1[0:10], 0)
	g2, e := strconv.Atoi(string(b1[0:index]))
	if e != nil {
		fmt.Println("atoi error")
	}
	fmt.Println("in data length=", n, "head length=", g2)
	if n > 10+g2 {
		s2, s3 := filepath.Split(string(b1[10 : g2+10]))
		fmt.Println(s2 + "  " + s3)
		os.Remove(s3)
		s1, _ = os.OpenFile(s3, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
		s1.Write(b1[10+g2 : n])

		for {
			n, e := t3.Read(b1)
			if e != nil {
				return
			}
			s1.Write(b1[0:n])
			fmt.Println(b1)
		}
	} else if n == 10+g2 {
		s2, s3 := filepath.Split(string(b1[10 : g2+10]))
		fmt.Println(s2 + "  creat file:" + s3)
		os.Remove(s3)
		s1, _ = os.OpenFile(s3, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
		for {

			n, e := t3.Read(b1)

			if e != nil {
				return
			}
			s1.Write(b1[0:n])
		}
	} else if n == 10 {
		n, e := t3.Read(b1)

		if e != nil {
			return
		}
		fmt.Println(b1)
		s2, s3 := filepath.Split(string(b1[0:g2]))
		fmt.Println(s2 + "  " + s3)
		os.Remove(s3)
		s1, _ = os.OpenFile(s3, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
		s1.Write(b1[g2:n])
		s1.Close()
		for {
			n, e := t3.Read(b1)
			if e != nil {
				return
			}
			fmt.Println(b1)
			s1, _ = os.OpenFile(s3, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
			s1.Write(b1[0:n])
			s1.Close()
			fmt.Println(b1)

		}

	} else {
		fmt.Println("package is len< wenjianname 10")
		return
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
