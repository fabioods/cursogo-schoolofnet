package function

func FuncName(a int) int {
	return a * a
}

func NamedReturn(a int, b int) (soma int) {
	soma = a + b
	return
}

func MoreReturn(a int, b int) (soma int, sub int) {
	soma = a + b
	sub = a - b
	return
}

func VariadicFunc(x ...int) int {
	res := 0
	for _, value := range x {
		res += value
	}
	return res
}

func FuncInsideFunc() func() int {
	x := 10
	otherFunction := func() int {
		return x * x
	}
	return otherFunction
}
