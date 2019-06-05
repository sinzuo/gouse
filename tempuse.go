// tempuse
package main

import (
	"fmt"

	"os"
	"strings"

	"text/template"
)

func returnBool(b bool) bool {
	return b
}

var s2 template.FuncMap = make(map[string]interface{})

func Add(in1 int, in2 int) int {
	return 2
}

var ss = `{{if is_teacher_coming .}}Carefully!{{end}}`

func main() {
	s1, _ := os.OpenFile("A.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	defer s1.Close()
	strings.NewReader(ss)

	l1 := template.New("test")
	s2["add"] = Add
	s2["is_teacher_coming"] = returnBool
	l1.Funcs(s2)
	data1 := true

	template.Must(l1.Parse(ss))
	l1.Execute(s1, data1)

	fmt.Println("Hello World!")
}
