package main

import (
	"fmt"
	"songo/blankidentifier"
	"songo/condicional"
	"songo/function"
	"songo/hello"
	"songo/lacos"
	"songo/pointers"
	"songo/switchcase"
	"songo/variables"
	"songo/visibility"
)

func main() {
	fmt.Printf("Hello, World! %v \n\n", z)
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
