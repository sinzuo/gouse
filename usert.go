package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var logok *log.Logger

func init() {
	logok = log.New(os.Stdout, "", log.LstdFlags)
}

func main() {

	jiang := strings.Fields("jiang yi bo ")
	ss, err := ioutil.ReadFile("/work/gouse/jiang.txt")
	if err != nil {
		fmt.Println("oo")
	}
	logok.Println("mmm")
	fmt.Println("jiangyibo" + jiang[2] + string(ss))
}
