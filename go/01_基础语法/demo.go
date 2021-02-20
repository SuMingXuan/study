// 定义包名
package main

import "fmt"
import "unsafe"
// 疑问: 这里只能引入文件夹吗？ 能不能只引入一个文件？
// 回答: 一个文件夹下面所有的文件定义的package都应该必须一致，否则就会报错，怪不得只能映入文件夹
import "./math"
import "./variable_test"
func main()  {
	fmt.Println("hello world")
	fmt.Println(myMath.Add(1,2))
	fmt.Println(myMath.Sub(2,1))
	fmt.Println(variableTest.Output_string("test"))
	fmt.Println(variableTest.Output_int())
	fmt.Println(variableTest.Output_error())
	// 省略 类型定义
	var a = 1
	var b = "1"
	var c = true
	var d = 0.1
	fmt.Println(a, b, c, d)
	// 省略 var(应该是常用的方式, 只能在函数中出现)
	e := 1
	fmt.Println(e)
	f, g := 0.1, true
	var h = g
	fmt.Println(f, g, h)
	const i int = 10
	const j, k = 1.01, "str"
	fmt.Println(i, j, k)
	const (
		l = "hello"
		m = len(l)
		n = unsafe.Sizeof(l)
		o = iota
		p = iota
	)
	const (
		q = iota
	)
	// iota 是个特殊的常量，在同一个const中会自增1， 目前还想不出会有哪些应用场景
	fmt.Println(l, m, n)
	fmt.Println(o, p, q)
}
