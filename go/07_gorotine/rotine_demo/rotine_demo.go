package rotine_demo
import (
	"fmt"
	"time"
	"sync"
)

func Loop()  {
	for i := 0; i < 11; i++ {
		time.Sleep(time.Microsecond * 1)

		fmt.Printf("%d,", i)
	}
}

var chanInt chan int = make(chan int, 2)
var timeOut chan bool = make(chan bool)

func Send()  {
	chanInt <- 1
	chanInt <- 2
	chanInt <- 3
	timeOut <- true
	timeOut <- true
}

func Receive()  {
	// num := <- chanInt
	// fmt.Println("num:", num)
	// num = <- chanInt
	// fmt.Println("num:", num)
	// num = <- chanInt
	// fmt.Println("num:", num)

	// select 配合chan使用，使其不同的chan做不同的事
	for {
		select {
		case num := <- chanInt:
			fmt.Println("num:", num)
		case <- timeOut:
			fmt.Println("timeout")

		}
	}
}


var WG sync.WaitGroup
func Read()  {
	WG.Add(5)
	// for i := 0; i < 5; i++ {
	// 	time.Sleep(time.Second * 2)
	//  WG.Add(1)
	// 	fmt.Println("Read:", i)
	// }
}

func Write()  {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 1)
		WG.Done()
		fmt.Println("Write:", i)
	}
}