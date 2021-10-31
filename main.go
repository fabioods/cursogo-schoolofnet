package main

import (
	"cursogo-schoolofnet/blankidentifier"
	"cursogo-schoolofnet/condicional"
	"cursogo-schoolofnet/function"
	"cursogo-schoolofnet/hello"
	"cursogo-schoolofnet/lacos"
	"cursogo-schoolofnet/pointers"
	"cursogo-schoolofnet/switchcase"
	"cursogo-schoolofnet/variables"
	"cursogo-schoolofnet/visibility"
	"fmt"
)

func main() {
	a := 1
	fmt.Printf("Hello, World! %v \n\n", a)
	hello.ShowUUID()
	fmt.Printf("Varible B %v\n\n", variables.B)
	variables.ShowVariables()
	visibility.PrintY()
	visibility.PrintX()
	blankidentifier.Blank()
	pointers.Show()
	condicional.Show()
	lacos.Show()
	switchcase.Show()
	fmt.Println("\n\nFunçoes")
	fmt.Println(function.FuncName(1))
	fmt.Println(function.NamedReturn(2, 3))
	fmt.Println(function.MoreReturn(4, 5))
	fmt.Println(function.VariadicFunc(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	z := 2
	add := func() int {
		z += 1
		return z
	}
	fmt.Printf("Z = %v\n\n", add())
	fmt.Printf("Z = %v\n\n", z)
	fmt.Println(function.FuncInsideFunc()())
}
