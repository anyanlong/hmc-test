package main

import (
	"fmt"
	"os"
)

var pidFd *os.File

func main() {
	var err error
	if pidFd, err = os.Open("/Users/hmc/work/gowork/src/hmc-test/file/complex_open.go"); err != nil {
		fmt.Println("open file err =", err)
	}
	bt := make([]byte, 1024)
	pidFd.Read(bt)
	fmt.Printf("file= %v", string(bt))

}
