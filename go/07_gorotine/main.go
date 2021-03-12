package main
import (
	"./rotine_demo"	
	"time"
	"runtime"
	"fmt"
)
func main()  {
	// runtime.NumCPU() 获取系统逻辑核心数
	// runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	runtime.GOMAXPROCS(3)
	// go rotine_demo.Loop()
	// go rotine_demo.Loop()

	// 协程通信
	// go rotine_demo.Send()
	// go rotine_demo.Receive()

	// 协程同步
	rotine_demo.Read()
	go rotine_demo.Write()
	rotine_demo.WG.Wait()
	fmt.Println("All done")

	time.Sleep(time.Second * 1)

}