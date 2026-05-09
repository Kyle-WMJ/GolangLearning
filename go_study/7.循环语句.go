package main

import (
	"fmt"
)

var sum int

func main() {

	// for循环
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 无限循环
	// for true {
	// 	fmt.Println(time.Now())
	// 	time.Sleep(2 * time.Second)
	// }

	// golang没有while循环,但是可以使用for循环模拟while循环
	var sum int
	var i int = 1
	for i <= 100 {
		sum += i
		i++
	}
	fmt.Println(sum)
	sum = 0
	i = 1
	for {
		sum += i
		i++
		if i > 100 {
			break
		}
	}
	fmt.Println(sum)

	// 循环切片或数组
	var list = []string{"a", "b", "c"}
	for i = 0; i < len(list); i++ {
		fmt.Println(list[i])
	}

	for index, item := range list {
		fmt.Println(index, item)
	}

	// 循环map
	var usermap = map[string]string{"name": "枫枫", "age": "18"}
	for k, v := range usermap {
		fmt.Println(k, v)
	}

	// 打印乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d ", i, j, i*j)
		}
		fmt.Println()
	}
}
