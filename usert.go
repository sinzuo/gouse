package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"os"
	"reflect"
	"strings"

	"time"
	"unsafe"

	"github.com/bitly/go-simplejson"
)

var logok *log.Logger

func init() {
	logok = log.New(os.Stdout, "", log.LstdFlags)
}

type A struct {
	Bt int
}

func (a *A) Str() {
	fmt.Println("ok")
}

const (
	ss int = iota
	mm
)

type D interface {
	Open()
}

func oo() {
	m1 := make(chan int)
	for i := 0; i < 10; i++ {
		go func() { m1 <- 1 }()
	}
	for i := 0; i < 10; i++ {
		s3 := <-m1
		fmt.Println(s3)
	}
	jiang := strings.Fields("jiang yi bo ")
	ccd := reflect.TypeOf(ss)
	var q1 int = 1
	qwe := reflect.ValueOf(&q1)
	qwe.Elem().SetInt(2)
	s8 := uintptr(unsafe.Pointer(&q1))
	s9 := (*int)(unsafe.Pointer(&q1))
	*(*int)(unsafe.Pointer(s8)) = 4
	fmt.Println(*s9)
	fmt.Println(ccd.Name() + "jjjj" + reflect.TypeOf(qwe).String() + os.Args[0] + time.Now().Format("2017:12:1"))
	sm, _ := simplejson.NewJson([]byte("{\"jinag\":\"bo\"}"))

	ss, err := ioutil.ReadFile("/work/gouse/jiang.txt")
	if err != nil {

		fmt.Println("oo")
		defer fmt.Println("rrr")
		return

	}

	oooo, _ := sm.String()
	logok.Println("mmm" + oooo)

	fmt.Println("jiangyibo" + jiang[2] + string(ss))
}
