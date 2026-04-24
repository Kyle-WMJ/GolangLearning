package main

import "fmt"

func main() {
	// 字典 key-value映射, 创建必须赋值或初始化，否则会报错
	// map1 := make(map[string]int)
	// fmt.Println(map1)
	var userMap map[int]string = map[int]string{
		1: "WMJ",
		2: "张三",
		3: "李四",
	}
	fmt.Println(userMap)
	fmt.Println(userMap[1])
	fmt.Printf("%#v\n", userMap[4]) // 不存在的key, 不报错，返回默认值

	value := userMap[3] // 单个赋值
	fmt.Println(value)
	value, ok := userMap[3] // 双赋值，ok时bool，表示有没有这个key
	fmt.Println(value, ok)

	delete(userMap, 3) // 删除key为3的元素
	fmt.Println(userMap)

	// 会报错，因为aMap没有赋值，所以不能直接赋值给它
	// var aMap map[string]string
	// aMap["age"] = "15"
	// fmt.Println(aMap)

	// 正确写法
	var aMap = map[string]string{}
	aMap["age"] = "15"
	fmt.Println(aMap)

	var bMap = make(map[string]string)
	bMap["gender"] = "male"
}
