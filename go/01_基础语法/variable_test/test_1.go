package variableTest

func Output_string(str string) string {
	var a string
	a = "这是变量a"
	var b string = "这是变量b"
	var c, d string = "这是变量c", "以及d"
	var e string
	return "输入了" + str + a + b + c + d + e 
}

func Output_int() int {
	var a int
	return a
}

func Output_error() error {
	var a error
	return a
}
