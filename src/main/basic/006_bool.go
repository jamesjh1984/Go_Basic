package main

import (
	"fmt"
	"unsafe"
)


func main() {

	var a = false
	fmt.Println("a =", a)
	fmt.Printf("a type = %T, a size = %d \n", a, unsafe.Sizeof(a))
}