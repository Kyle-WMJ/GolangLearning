package main

import (
	"fmt"
)

func main() {
	// 输入
	fmt.Println("hello world")
	fmt.Println("你好", "WMJ")        // 自带空格隔开，换行
	fmt.Printf("哇！你好美！%s\n", "GYY") // 格式化输出
	fmt.Printf("哇！你好美！%s\n", "WMJ")
	fmt.Printf("整数 %d\n", 1)
	fmt.Printf("浮点数 %.2f\n", 3.1415926)
	fmt.Printf("类型 %T %T\n", "3.1415926", 3.1415926)
	fmt.Printf("任意类型 %v\n", "3.1415926")
	fmt.Printf("加井号可打印出空字符串 %#v\n", "")

	// 把输出赋给一个变量,Sprintf返回一个格式化字符串
	var f = fmt.Sprintf("%.2f", 1.414)
	fmt.Println(f)

	// 输出
	fmt.Print("请输入你的名字：")
	var name string
	fmt.Scan(&name)
	// n, err := fmt.Scan(&name)
	// if err != nil {
	//	fmt.Println(n, err)
	// }
	fmt.Println(name)
	fmt.Print("请输入你的年龄：")
	var age int
	fmt.Scan(&age)
	fmt.Println(age)

}
