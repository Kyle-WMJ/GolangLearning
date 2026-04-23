package main

import "fmt"

func main() {
	/*
		整数类型
		uint8 无符号8位整数
		int8 有符号8位整数
		uint16 无符号16位整数
		int16 有符号16位整数
		uint32 无符号32位整数
		int32 有符号32位整数
		uint64 无符号64位整数
		int64 有符号64位整数
		uint 无符号整数, 根据系统位数不同而不同
		int 有符号整数, 根据系统位数不同而不同
	*/
	fmt.Println("---------整数类型---------")
	var u8 uint8 = 255
	fmt.Println(u8)

	/*
		浮点数类型
		float32 32位浮点数
		float64 64位浮点数
		float 默认64位浮点数
	*/
	fmt.Println("---------浮点数类型---------")
	var f32 float32 = 3.1415926
	var f64 float64 = 3.1415926
	fmt.Println(f32, f64)

	/*
		字符类型
		rune 多字节字符，中文，韩文等
		byte 单字节字符，ascii里面的字符
	*/
	fmt.Println("---------字符类型---------")
	var a byte = 'a'
	fmt.Printf("%c %d\n", a, a)
	var a1 uint8 = 97
	fmt.Printf("%c %d\n", a1, a1)

	var r rune = '中'
	fmt.Printf("%c %d\n", r, r)
	var r1 int32 = 20013
	fmt.Printf("%c %d\n", r1, r1)

	/*
		string 字符串类型
		`多行
		字符
		串`
	*/
	fmt.Println("---------字符串类型---------")
	var s string = "我知道"
	fmt.Println(s)
	fmt.Println("'我'知道")
	fmt.Println("\"我\"知道")                         //转义字符 双引号
	fmt.Println("C:\\Windows\\System32\\notepad.exe") //转义字符 反斜杠
	fmt.Println("枫枫枫枫\r知道")
	fmt.Println(`hello
\n
world`) //反引号可以跨行，也可以嵌套，``里面的转义字符无效

	/*布尔类型
	bool 布尔值，只有true和false两个值
	无法参与计算，默认为false
	go语言不允许将整型强制转化为布尔型
	*/

	var b bool = true
	fmt.Println(b)

	/*
		零值问题
		若给一个基本数据类型只声明不赋值，会自动赋值为零值
		如int就是0，bool就是false
	*/
	var str string
	fmt.Printf("%#v\n", str) //#v 打印出数据原始值
	var cha byte
	fmt.Printf("%#v\n", cha)
	var b1 bool
	fmt.Printf("%#v\n", b1)
}
