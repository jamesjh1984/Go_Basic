package main

import (
	"fmt"
	"unsafe"
)


func main() {

	var a float32 = 22.5
	fmt.Println("a =", a)

	// 默认数据类型(float64)，占用字节大小
	var b = 33.5
	fmt.Printf("b type = %T, b size = %d \n",b, unsafe.Sizeof(b))

	// 科学计数法
	c := 3.1415926e2 // 3.1415926 * 10的2次方
	d := 3.1415926E2 // 3.1415926 * 10的2次方
	e := 3.1415926E-2 // 3.1415926 / 10的2次方
	fmt.Println("c =", c, ",d =",d, ", e =",e)

}