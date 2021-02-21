package main

import "fmt"

func main()  {
	a := 11
	if a < 10 {
		fmt.Println("a 小于 10")
	} else {
		fmt.Println("a 大于 10")
	}

	for i := 0; i < a; i++ {
		fmt.Println(i)
	}
	b := 1
	fmt.Println("最大值是：", max(a, b))
	c, d := "str1", "str2"
	fmt.Println(swap(c, d))
	e := []int{1,2,3}
	fmt.Println(e)
	fmt.Println(swap1(c, d))
	fmt.Println(c, d)
	fmt.Println(swap2(&c, &d))
	fmt.Println(c, d)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 怎么定义不同类型的参数
func swap(x, y string) (string, string) {
	return y, x
}

// 值传递
func swap1(x, y string) (string, string) {
	var temp string
	temp = x
	x = y
	y = temp
	return x, y
}

// 引用传递
func swap2(x *string, y *string) (string, string) {
	var temp string
	temp = *x
	*x = *y
	*y = temp
	return *x, *y
}