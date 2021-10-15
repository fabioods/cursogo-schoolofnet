package switchcase

import "fmt"

func Show() {
	println("\n\nSwitch case")
	a := "Fábio0"
	switch a {
	case "Fábio":
		fmt.Print("Fábio")
	case "João":
		fmt.Print("João")
	default:
		fmt.Print("Outro")
	}
}
