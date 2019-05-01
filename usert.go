package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"os"
	"reflect"
	"strings"
	"time"

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

func main() {

	jiang := strings.Fields("jiang yi bo ")
	ccd := reflect.TypeOf(ss)
	reflect.ValueOf(1)
	fmt.Println(ccd.Name() + os.Args[0] + time.Now().Format("2017:12:1"))
	sm, _ := simplejson.NewJson([]byte("{\"jinag\":\"bo\"}"))

	ss, err := ioutil.ReadFile("/work/gouse/jiang.txt")
	if err != nil {

		fmt.Println("oo")
		defer fmt.Println("rrr")
		return

	}
     ii chan
	oooo, _ := sm.String()
	logok.Println("mmm" + oooo)

	fmt.Println("jiangyibo" + jiang[2] + string(ss))
}
