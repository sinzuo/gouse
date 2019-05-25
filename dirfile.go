// dirfile
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

func A(dir string) {
	f1, _ := ioutil.ReadDir(dir)
	for _, f := range f1 {
		if f.IsDir() {
			A(path.Join(dir, f.Name()))
		} else {
			fmt.Println(path.Join(dir, f.Name()))
		}
	}
}

func main() {
	s1 := filepath.Dir(os.Args[0])
	s2, s3 := path.Split(os.Args[0])
	A(s1)
	fmt.Println("Hello World!", s1, s2, s3)

}
