// 变量知识点
package main

import (
	"fmt"
	"reflect"
)
func main()  {
	// create variable
	// makeSlice()
	// makeMap()
	// newMap()
	// update variable
	// appendElementForSlice()
	// copyForSlice()
	deleteForMap()
}
// 创建变量 make, new

// make -----------------start------------------- 
// 返回的是引用类型
// slice map chan 都需要使用make来创建，能让go自动为我们初始化好，避免出现问题
func makeSlice() {
	// string 是数组里面的类型, 3是长度
	mSlice := make([]string, 3)
	mSlice[0] = "dog"
	mSlice[1] = "cat"
	mSlice[2] = "trigger"
	fmt.Println(mSlice)
}

func makeMap()  {
	// int 是key的类型 string 是value的类型
	mMap := make(map[int]string)
	mMap[5] = "s"
	mMap[7] = "m"
	mMap[9] = "x"
	mMap[10] = "x"
	fmt.Println(mMap)
}

func makeChan()  {
	// int 是传输的类型， 3 是缓存数量， 如果不要3的话就是一个没有缓存的chan
	mChan := make(chan int, 3)
	close(mChan)
}
// make -----------------end------------------- 


// new -----------------start------------------- 
// 特性
// 返回的是值类型
// 内存置零
// 返回传入类型的指针地址

func newMap()  {
	nMap := new(map[int]string)
	mMap := make(map[int]string)
	// 引用类型
	fmt.Println("make map:", reflect.TypeOf(mMap))
	// 指针类型
	fmt.Println("new map:", reflect.TypeOf(nMap))
}
// new -----------------end------------------- 



// 操作变量 append, delete, copy

// slice: append, delete
// map:   delete

func appendElementForSlice()  {
	mSlice := make([]string, 2)
	mSlice[0] = "ruby"
	mSlice[1] = "go"
	fmt.Println("before update len:", len(mSlice))
	fmt.Println("before update cap:", cap(mSlice))
	mSlice    = append(mSlice, "java")
	mSlice    = append(mSlice, "JS")
	mSlice    = append(mSlice, "Rust")
	fmt.Println(mSlice)
	fmt.Println("after update len:", len(mSlice))
	fmt.Println("after update cap:", cap(mSlice))
}

func copyForSlice()  {
	mSliceTarget := make([]string, 2)
	mSliceTarget[0] = "ruby"
	mSliceTarget[1] = "go"
	// mSliceTarget[2] = "java"

	mSliceSource := make([]string, 3)
	mSliceSource[0] = "dog"
	mSliceSource[1] = "cat"
	mSliceSource[2] = "trigger"
	fmt.Println("before copy target:", mSliceTarget)
	fmt.Println("before copy source:", mSliceSource)
	// 第二个复制到第一个去, 长度如果不一致的话，第一个的长度不会被改变
	copy(mSliceTarget, mSliceSource)
	fmt.Println("after copy target:", mSliceTarget)
	fmt.Println("after copy source:", mSliceSource)
}

func deleteForMap()  {
	mMap := make(map[string]int)
	mMap["ruby"] = 9
	mMap["java"] = 1
	mMap["linux"] = 7
	mMap["js"] = 7
	fmt.Println("before delete:", mMap)
	delete(mMap, "linux")
	fmt.Println("after delete:", mMap)
}
