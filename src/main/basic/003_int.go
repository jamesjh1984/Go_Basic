package main

import (
	"fmt"
	"unsafe"
)


func main() {

	var a int = 1
	fmt.Println("a =", a)


	// 测试int8的范围（-128 ~ 127）
	//var b int8 = -129
	var b int8 = 127
	fmt.Println("b(int8) =", b)


	// 测试uint8的范围（0 ~ 255）
	//var c uint8 = 256
	var c uint8 = 0
	fmt.Println("c(uint8) =", c)


	// 测试int的范围,64位（-2(63方) ~ 2(63方)-1）
	//var d int = -129
	var d int = 8900
	fmt.Println("d(int) =", d)


	// 测试uint的范围,64位（-2(63方) ~ 2(63方)-1）
	//var e uint = -129
	var e uint = 1
	fmt.Println("e(uint) =", e)


	// 默认数据类型(int)，占用字节大小
	var f = 100
	fmt.Printf("f type = %T, f size = %d \n", f, unsafe.Sizeof(f))


}