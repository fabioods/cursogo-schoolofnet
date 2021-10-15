package lacos

import "fmt"

func Show() {
	fmt.Println("\n\nfor")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}

	y := 0
	for {
		y++
		if y == 50 {
			continue
		}

		if y == 100 {
			break
		}
	}
}
