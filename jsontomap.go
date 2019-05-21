// jsontomap
package main

import (
	"encoding/json"
	"fmt"
)

var s1 = `{"dog":[{"sid":2,"tianya":"sfdsaf"}]}`

type A struct {
	Sidd   int
	Tianya string
}

type B struct {
	Dog []A
}

func main() {
	//	var sa = make(map[string]int, 10)
	sa := &B{}
	json.Unmarshal([]byte(s1), &sa)

	fmt.Println(sa.Dog[0].Tianya)

	fmt.Println("Hello World!", s1)
}
