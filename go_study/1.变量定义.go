// main才能执行
package main

import (
	"fmt"
	"go_study1/go_study/version"
)

func hello() {
	fmt.Println(age) //全局变量不遵守先后顺序，也不必一定要用
	fmt.Println("hello world")
}

var age = 12

const version1 = "2.0.1" //常量必须赋值

var (
	s1 string = "str1"
	s2 string = "str2"
)

func main() {
	// 变量定义后，必须使用赋值，否则会报错
	var name string
	name = "WMJ"
	fmt.Println(name)

	// 声明并赋值
	var name1 string = "WMJ1"
	fmt.Println(name1)

	//省略类型
	var name2 = "WMJ2"
	fmt.Println(name2)

	//声明并赋值 :声明符号
	name3 := "WMJ3"
	fmt.Println(name3)

	hello()

	var a1, a2, a3 = 1, 2, 3
	fmt.Println(a1, a2, a3)

	fmt.Println(s1, s2)

	fmt.Println(version.Version2)
	//fmt.Println(version.fengfeng)
}
