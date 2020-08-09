package curso1

import (
	"fmt"
)

// ExecuteGoCap12Ej1 hace tal cosa.
func ExecuteGoCap12Ej1() {
	fmt.Printf("\n\nExecuteGoCap12Ej1 <=== ****\n")
	var puntero1, puntero2 *int
	entero := 5

	puntero1 = &entero
	puntero2 = &entero

	fmt.Printf("Variable entero: %d\n", entero)
	fmt.Printf("Direccion de memoria de la variable entero: %d\n", &entero)
	fmt.Printf("Valor de puntero1: %d | Valor apuntado por puntero1: %d\n", puntero1, *puntero1)
	fmt.Printf("Valor de puntero2: %d | Valor apuntado por puntero2: %d\n", puntero2, *puntero2)
	fmt.Println()

	fmt.Println("Cambiando el valore por medio del puntero1")
	*puntero1 = 2
	fmt.Printf("Variable entero: %d\n", entero)
	fmt.Printf("Direccion de memoria de la variable entero: %d\n", &entero)
	fmt.Printf("Valor de puntero1: %d | Valor apuntado por puntero1: %d\n", puntero1, *puntero1)
	fmt.Printf("Valor de puntero2: %d | Valor apuntado por puntero2: %d\n", puntero2, *puntero2)
	fmt.Println()
}
