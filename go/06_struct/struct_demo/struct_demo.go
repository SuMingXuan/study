package struct_demo
// 定义 dog 的结构体
import "fmt"

type Animal struct {
	Id int
}

type Dog struct {
	Animal
	Name string
	Age int
}

type Cat struct {
	Animal
	Name string
	Age int
}

func TestForStructFirst()  {
	var dog Dog
	dog.Id = 1
	dog.Name = "xiaobai"
	dog.Age = 18
	fmt.Println("dog: ", dog)
}

func TestForStructSecond()  {
	// 这种方式不能使用id
	dog := Dog{Name: "xiaohei", Age: 3}
	fmt.Println("dog: ", dog)
}

func TestForStructThird()  {
	// 值类型， 指针
	dog := new(Dog)
	dog.Id = 3
	dog.Name = "xiaohong"
	dog.Age = 14
	fmt.Println("dog: ", dog)
}

func (d *Dog)Run() string {
	fmt.Println("Id: ", d.Id, "dog runing")
	return "dog runing"
}

func (a *Animal)Eat() string {
	fmt.Println("Id: ", a.Id, "animal eating")
	return "animal eating"
}

func (d *Cat)Run() string {
	fmt.Println("Id: ", d.Id, "cat runing")
	return "cat runing"
}

