package main
import "./interface_demo"
import "./struct_demo"
func main()  {
	// struct_demo.TestForStructFirst()
	// struct_demo.TestForStructSecond()
	// struct_demo.TestForStructThird()

	// dog1 := struct_demo.Dog{Name: "xiaohei", Age: 3}
	// dog1.Run()

	// dog2 := new(struct_demo.Dog)
	// dog2.Id = 4
	// dog2.Name = "xiaohuan"
	// dog2.Age = 14
	// dog2.Run()
	// dog2.Eat()

	// 接口测试
	dog := new(struct_demo.Dog)
	cat := new(struct_demo.Cat)
	action(dog)
	action(cat)
}

func action(b interface_demo.Behavior) string {
	b.Run()
	b.Eat()
	return "over"
}