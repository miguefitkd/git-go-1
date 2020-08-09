package curso1

import (
	"fmt"
)

// ExecuteGoCap6Ej1 hace tal cosa
func ExecuteGoCap6Ej1() {
	fmt.Printf("\n\nExecuteGoCap6Ej1 <=== ****\n")
	x := 10
	y := 5

	if x > y {
		fmt.Printf("%d es mayor que %d\n", x, y)
	} else if x < y {
		fmt.Printf("%d es menor que %d\n", x, y)
	} else {
		fmt.Printf("Los valores son iguales.\n")
	}
}
