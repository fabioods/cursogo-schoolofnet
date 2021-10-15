package condicional

import "fmt"

func Show() {
	a := 5
	fmt.Print("\n\nUso de condicional\n\n")
	if a > 10 {
		fmt.Println("a > 10")
	} else if a > 5 {
		fmt.Print("a > 5")
	} else {
		fmt.Print("a < 5")
	}

	if x := 10; x > 5 {
		fmt.Println("x > 5")
	} else {
		fmt.Println("x < 5")
	}
}
