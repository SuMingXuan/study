// 异常知识点
// panic 抛出异常
// recover 捕获异常
package main
import (
	"fmt"
	"errors"
)
func main()  {
	receivePanic()
}

func receivePanic()  {
	defer rescue()
	// panic("i'm panic")
	// panic(1)
	panic(errors.New("i'm error type"))
}

func rescue()  {
	message := recover()
	switch message.(type) {
	case string:
		fmt.Println("string: ", message)
	case error:
		fmt.Println("error: ", message)
	default:
		fmt.Println("unknown: ", message)
	}
}