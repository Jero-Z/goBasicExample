package main

import "fmt"

//变量的类型和声明

func main() {
	// 字符串类型必须使用双引号
	var a = "variable"
	fmt.Println(a)
	// 可批量声明变量
	var b, c int = 1, 2
	fmt.Println(b, c)
	// 声明布尔值类型变量
	var d = true
	fmt.Println(d)
	// int变量默认值为0
	var e int
	fmt.Println(e)
	// string 默认为"" 长度为0
	var f string
	fmt.Println(f, len(f))

}
