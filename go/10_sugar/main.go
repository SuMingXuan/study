package main
import "fmt"
func main()  {
	sugar("a", "b", "c", "d", true, make(map[string]int))
}
// 不限个数以及类型的参数，
func sugar(values...interface{})  {
	// for index, value := range values {
	// 	fmt.Println("--->", index, value)
	// }
	fmt.Println("--->", values[2])
}