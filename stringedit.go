// stringedit
package main

import (
	"fmt"
	//	"io"
	"strings"
	"unicode/utf8"
)

var s1 = `"是字符串 %s "`

func main() {

	s := fmt.Sprintf(s1, "string")
	s2 := strings.ToLower("jiang  sf")
	utf8.RuneCountInString("jiang")
	fmt.Println(s, s2) // 是字符串 %s  对应  是字符串 string
}
