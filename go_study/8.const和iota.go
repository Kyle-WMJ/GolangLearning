package main

import "fmt"

// 可以在const()添加一个关键字iota，每行的iota都会累加1，第一行的iota默认为0
const (
	Beijing = iota * 10
	Shanghai
	Guangzhou
)

var city []int = []int{Beijing, Shanghai, Guangzhou}

func main() {
	for _, name := range city {
		fmt.Println(name)
	}
	var city2 []int = make([]int, 0, 0)
	city2 = append(city2, Beijing, Shanghai, Guangzhou)
	fmt.Println(city2)

	const (
		a, b = iota + 1, iota + 2
		c, d
		e, f

		g, h = iota * 10, iota * 20
		i, j
	)
	fmt.Println("a= ", a, " b= ", b)
	fmt.Println("c= ", c, " d= ", d)
	fmt.Println("e= ", e, " f= ", f)
	fmt.Println("g= ", g, " h= ", h)
	fmt.Println("i= ", i, " j= ", j)
}
