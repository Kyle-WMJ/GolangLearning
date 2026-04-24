package main

import "fmt"

func main() {
	// 字典 key-value映射
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
}
