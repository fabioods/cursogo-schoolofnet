package pointers

import "fmt"

func soma(valor *int) int {
	*valor = *valor + 10
	return *valor
}

func Show() {
	fmt.Printf("\n\n\nusando ponteiros")
	x := 10
	fmt.Printf("\nvalor de x: %d", x)
	fmt.Printf("\nendereço de x: %v", &x)

	y := &x
	fmt.Printf("\nendereço de y: %v", y)
	fmt.Printf("\nvalor de y: %v", *y)

	*y = 20
	fmt.Printf("\nendereço de y: %v", y)
	fmt.Printf("\nvalor de y: %v", *y)
	fmt.Printf("\nvalor de x: %v", x)

	var z *int = &x
	fmt.Printf("\nendereço de z: %v", z)
	fmt.Printf("\nvalor de z: %v", *z)
	fmt.Print("\n\n")

	a := 10
	fmt.Printf("soma: %v, a = %v", soma(&a), a)

	fmt.Print("\n")
}
