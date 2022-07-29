package main

import (
	"fmt"
	"unsafe"
)


func main() {

	var a string = "Hangzhou"
	fmt.Println(a)
	fmt.Printf("a type = %T, a size = %d \n", a, unsafe.Sizeof(a))


	// string赋值后就不可以修改
	//var b string = "Hangzhou"
	//b[0] = 'h'


	// 特殊字符
	var c string = "Zhejiang\nHangzhou\n"
	fmt.Println(c)


	// 反引号``原生输出
	var d string = `package main
import "fmt"
func main() {
	fmt.Println("Hello Go!");	fmt.Println("Hello Go!")
}`
	fmt.Println(d)



	// 字符串拼接
	var e string = "hello " + "world! "
	e += "welcome"
	fmt.Println(e)

	// 当一个拼接很长时，可以分行写，但必须把+放在上一行尾
	var f string = "hello " + "world! " + "hello " + "world! " +
					"hello " + "world! " + "hello " + "world! " +
					"hello " + "world! " + "hello " + "world! " +
					"hello " + "world! " + "hello " + "world! " +
					"hello " + "world! " + "hello " + "world! "
	fmt.Println(f)


}