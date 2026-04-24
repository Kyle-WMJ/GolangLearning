package main

import (
	"fmt"
)

func main() {
	fmt.Println("请输出你的年龄：")
	var age int16
	_, err := fmt.Scanf("%d", &age)
	if err != nil {
		fmt.Println("输入错误，请输入数字", err)
		return
	}
	// 中断式，卫语句，逻辑清晰，推荐使用
	if age <= 0 {
		fmt.Println("未出生")
		return
	}
	if age <= 18 {
		fmt.Println("未成年")
	}
	if age >= 18 {
		fmt.Println("成年人")
	}

	// 嵌套判断语句
	if age >= 30 {
		if age <= 50 {
			fmt.Println("中年人")
		} else {
			fmt.Println("中老年人")
		}
	} else {
		if age >= 18 {
			fmt.Println("青年人")
		}
	}

	if age >= 100 {
		fmt.Println("长寿")
	} else if age <= 100 {
		fmt.Println("高龄")
	}

	// 多条件写法 &&与 ||或 !=非
	if age >= 100 && age <= 120 {
		fmt.Println("长寿")
	}

	// switch
	switch {
	case (age <= 18):
		fmt.Println("未成年")
	case (age >= 18):
		fmt.Println("成年人")
	}

	switch day := 1; day {
	case 1:
		fmt.Println("one")
		fallthrough // 强制继续执行下一个case语句
		// fmt.Println("three")
	case 2:
		fmt.Println("two")
	}

}
