package point_demo
import "fmt"
func TestPoint()  {
	var count int = 20
	var countPoint *int
	countPoint = &count
	fmt.Println("count 地址: \n", &count)
	fmt.Println("countPoint 地址: \n", countPoint)
	fmt.Println("countPoint 指针指向地址的值:", *countPoint)

}

func TestPointArr()  {
	// 指针数组
	a, b := 1, 2
	// * 代表他是一个指针， 指向int类型的指针
	pointArr := [...]*int{&a, &b}
	fmt.Println(pointArr)

	// 数组指针
	arr := [...]int{1, 2, 3}
	arrPoint := &arr
	fmt.Println(arrPoint)

}