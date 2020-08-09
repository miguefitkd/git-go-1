package curso1

import (
	"fmt"
)

// ExecuteGoCap7Ej1 hace tal cosa.
func ExecuteGoCap7Ej1() {
	fmt.Printf("\n\nExecuteGoCap7Ej1 <=== ****\n")

	for i := 1; i <= 10; i++ {
		fmt.Printf("for - Hola Mundo! %d\n", i)
	}

	fmt.Println("")

	y := 1
	for y <= 10 {
		fmt.Printf("for/while - Hola Mundo! %d\n", y)
		y++
	}
}
