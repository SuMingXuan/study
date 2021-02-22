package main

import "fmt"
func main()  {
	var arr1 [10] int
	arr1 = [10]int{1, 2, 4}
	fmt.Println(arr1)

	arr2 := [5]float32{1.0, 2.23, 1.2}
	fmt.Println(arr2)
	arr3 := [10]string{"a", "b"}
	fmt.Println(arr3)
	// 不确定数组多长
	arr4 := [...]int{1,2,3}
	fmt.Println(arr4)
	fmt.Println(arr4[2])
	// 初始化索引为1和4的元素， 元素6 的索引紧跟索引1的元素后面
	arr5 := [5]int{1:4, 6, 4:8}
	arr6 := [5]int{4, 6, 8}
	fmt.Println(arr5)
	fmt.Println(arr6)
	var arr7 [2][3][4]int
	fmt.Println(arr7)
	arr8 := [][]int{}
	fmt.Println(arr8)
	row1 := []int{1,2,3}
	row2 := []int{4,5,6,7}
	arr8 = append(arr8, row1)
	fmt.Println(arr8)
	arr8 = append(arr8, row2)
	fmt.Println(arr8)
	arr9 := [][]int{{1,2,3}, {4,5,6,7,8}}
	fmt.Println(arr9)
	fmt.Println(arr9[1][3])
	arr10 := [3][4]int{
		{1, 2, 3, 4},
		{2, 2, 3, 4},
		{3, 2, 3, 4},
	}
	fmt.Println(arr10)
	var avg float32
	avg = getAverage(arr1, 10)
	fmt.Println(avg)
	fmt.Println(arr4)
	var slice1 []int
	var slice2 []int = make([]int, 7)
	slice3 := make([]int, 7)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(row1)
	slice4 := [7]int{1,2,3,4,5,6,7}
	fmt.Println(slice4[0:2])
	fmt.Println(slice4[2:])
	fmt.Println(slice4[:5])
	
}

func getAverage(arr [10]int, size int) float32 {
	var i, total int
	for i = 0; i < size; i++ {
		total += arr[i]
	}
	return float32(total) / float32(size)
}

