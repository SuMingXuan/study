package interface_demo

type Behavior interface {
	// string 是返回类型，需要和方法的返回类型保持一致
	Run() string
	Eat() string
}