// syncwait
package main

import (
	"fmt"
	"sync"
	"time"
)

func ss(i int) {
	s1.Add(1)
	defer s1.Add(-1)
	fmt.Println(i)
	time.Sleep(time.Second * 5)
}

var s1 = &sync.WaitGroup{}

func main() {
	for i := 0; i < 5; i++ {

		go ss(i)
	}
	fmt.Println("wok")
	time.Sleep(time.Second * 1)
	fmt.Println("wok1")
	s1.Wait()
	fmt.Println("Hello World!")
}
