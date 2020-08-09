package curso1

import (
	"fmt"
	"strconv"
)

// ExecuteGoCap4Ej1 hace tal cosa.
func ExecuteGoCap4Ej1() {
	fmt.Printf("\n\nExecuteGoCap4Ej1() <=== ****\n")

	edad := 36
	fmt.Println(edad)

	// (type untyped string) as type int
	// fmt.Println("Mi edad es: " + edad)

	edadStr := strconv.Itoa(edad)
	fmt.Println("Mi edad es: " + edadStr)

	// El segundo valor es el error.
	edadInt, _ := strconv.Atoi(edadStr)
	fmt.Println(edadInt + 10)
}
