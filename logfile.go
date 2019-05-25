// logfile
package main

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type A struct {
	Length int
	Maxlen int
	Muxio  sync.Mutex
	Fileok io.Writer
}

func (a *A) Init() {
	fmt.Println("ok")
	f1, _ := os.OpenFile(time.Now().String(), os.O_APPEND|os.O_CREATE|os.O_RDWR|os.O_SYNC, os.ModePerm)
	a.Length = 0
	a.Maxlen = 15
	a.Fileok = f1
}

func (a *A) Write(b []byte) {
	a.Muxio.Lock()
	defer a.Muxio.Unlock()
	ba1 := len(b)
	if ba1+a.Length <= a.Maxlen {
		a.Fileok.Write(b)
		a.Length = a.Length + ba1
	} else {
		f1, _ := os.OpenFile(time.Now().String()+".txt", os.O_APPEND|os.O_CREATE|os.O_RDWR|os.O_SYNC, os.ModePerm)
		a.Fileok = f1
		a.Fileok.Write(b)
		a.Length = ba1

	}
}
func main() {
	mm := &A{}
	mm.Init()
	mm.Write([]byte("jiangyibo"))
	mm.Write([]byte("jiangyibo"))
	fmt.Println("Hello World!")
}
