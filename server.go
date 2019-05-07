// server
package main

import (
	"fmt"
	"net/http"
)

func handle(w http.ResponseWriter, req *http.Request) {

	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("jiangyibo"))

}

type 

//var s1 = http.Handle("bobo", handle)

func main() {
	

	http.HandleFunc("/jiang", handle)
	http.ListenAndServe(":9999", nil)
	fmt.Println("Hello World!")
}
