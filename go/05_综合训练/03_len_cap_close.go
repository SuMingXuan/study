// len, cap, close
// len: string array slice map chan
// cap: slice array chan
// close: chan
package main
import "fmt"
func main() {
	// getLenAndCap()
	closeChan()
}

func getLenAndCap()  {
	lenStr := "asd, qwe, zxc" 
	fmt.Println("lenStr len: ", len(lenStr))
	fmt.Println("lenStr cap: ", len(lenStr))
	// 设置容量
	lenSlice := make([]int, 3, 4)
	lenSlice[0] = 10
	lenSlice[1] = 20
	lenSlice[2] = 30
	lenSlice    = append(lenSlice, 40)
	lenSlice    = append(lenSlice, 50)
	fmt.Println("lenSlice len: ", len(lenSlice))
	fmt.Println("lenSlice cap: ", cap(lenSlice))
}

func closeChan()  {
	cChan := make(chan int, 1)
	// 相当于 ruby 的 ensure
	defer close(cChan)
	cChan <- 0
}