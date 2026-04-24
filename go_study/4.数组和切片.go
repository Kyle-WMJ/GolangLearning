package main

import (
	"fmt"
	"sort"
)

func main() {
	// 数组, 长度固定
	var nameList [3]string = [3]string{"张三", "李四", "王五"}
	fmt.Println(nameList, nameList[0], len(nameList))

	// 切片(列表), 长度可变
	var nameList2 []string = []string{"张三", "李四", "王五"}
	// var nameList2 = []string[]
	// var nameList2 := make([]string, 0)
	fmt.Println(nameList2, nameList2[0], len(nameList2))

	// 添加元素
	nameList2 = append(nameList2, "赵六")
	fmt.Println(nameList2, nameList2[0], len(nameList2))

	// 删除元素
	nameList2 = append(nameList2[:0], nameList2[1:]...)
	fmt.Println(nameList2, nameList2[0], len(nameList2))

	// 通过make函数创建指定长度，容量的切片, make([]类型, 长度, 容量)
	nameList3 := make([]string, 3, 4)
	fmt.Println(nameList3)
	ageList := make([]int, 3, 4)
	fmt.Println(ageList)
	nameList3 = append(nameList3[:0], []string{"张三", "李四", "王五", "赵六", "王二"}...)
	fmt.Println(nameList3, nameList3[0], len(nameList3), cap(nameList3))

	array := [3]int{1, 2, 3}
	slices := array[:]
	fmt.Println(slices)
	fmt.Println(array[0:2])
	fmt.Println(array[1:2]) // 于python的切片相同，包左不包右

	var ints = []int{4, 8, 3, 2}
	sort.Ints(ints) // 默认升序
	fmt.Println(ints)
	sort.Sort(sort.Reverse(sort.IntSlice(ints))) // 降序
	fmt.Println(ints)

}
