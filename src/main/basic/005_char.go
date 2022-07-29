package main

import (
	"fmt"
)


func main() {

	var a byte = 'a'
	var b byte = '0' // 字符的0，不是数字的0

	// 当我们直接输出byte值，就是输出了字符对应的ASCII码值
	fmt.Println("a =", a) // 97
	fmt.Println("b =", b) // 48

	// 如果希望输出对应字符，需要格式化输出
	fmt.Printf("a = %c, b = %c \n", a,b)


	//var c byte = '金' // 对应unicode码值为37329，超出byte范围(0-255)
	var c int = '中' // 可以改用int声明，20013
	var d int = '国' // 可以改用int声明，22269
	fmt.Printf("c = %c, c unicode = %d \n", c, c)
	fmt.Printf("d = %c, d unicode = %d \n", d, d)


	var e int = 20013 // 可以直接赋整数，然后格式化输出
	fmt.Printf("e = %c, e unicode = %d \n", e, e)
}