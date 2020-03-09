package main

import "fmt"

func main() {
	a := "123456"
	b:= a[1:5]
	a = "654321"
	fmt.Println("a= ", a, "\tb=", b)
}
