package variables

import "fmt"

var B int = 22
var hello, world string = "Hello", "World"

const constanteExterna = 11
const (
	constante1 = "iota"
	constante2 = "bitcoin"
)

func ShowVariables() {
	a := 1
	const constante = 10
	fmt.Printf("A %v B %v\n", a, B)
	fmt.Printf("hello (%v) world (%v)\n\n", hello, world)
	fmt.Printf("constante %v \n\n", constante)
	// constante = 11
	fmt.Printf("constanteExterna %v \n\n", constanteExterna)
	fmt.Printf("constante1 %v \n\n", constante1)
	fmt.Printf("constante2 %v \n\n", constante2)
}
