package main

import "fmt"

// 定义全局变量
// var g1, g2, g3 = 100, "James Jin", "Boss"
// or
var (
	g1 = 100
	g2 = "James Jin"
	g3 = "Boss"
)



func main() {

	// 第一种，声明+赋值
	var a int // 默认值0
	a = 10
	fmt.Println("a =", a)

	// 第二种，根据值自行判断变量类型
	var b = 10 // 默认值0
	fmt.Println("b =", b)

	// 第三种，省略var，注意:=左边的变量不能是一件声明过的，否则会编译错误
	c := 10
	fmt.Println("c =", c)

	// 第四种，一次声明多个变量
	var a1, a2, a3 int
	fmt.Println("a1 =",a1, ",a2 =",a2, ",a3 =",a3)

	// 第四种，一次声明多个变量, 根据值自行判断变量类型
	var b1, b2, b3 = 100, "James", "boy"
	fmt.Println("b1 =",b1, ",b2 =",b2, ",b3 =",b3)

	// 第四种，一次声明多个变量, 根据值自行判断变量类型
	c1, c2, c3 := 100, "James", "boy"
	fmt.Println("c1 =",c1, ",c2 =",c2, ",c3 =",c3)





	fmt.Println("--- 全局变量 ----")
	fmt.Println("g1 =",g1, ",g2 =",g2, ",g3 =",g3)


}