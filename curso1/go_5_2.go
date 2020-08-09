package curso1

import (
	"bufio"
	"fmt"
	"os"
)

// ExecuteGoCap5Ej2 - Leer e imprimir datos.
func ExecuteGoCap5Ej2() {
	fmt.Printf("\n\nExecuteGoCap5Ej2 <=== ****\n")
	// Hay que indicarle que lector tiene que utilizar
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingresa tu nombre: ")
	// Definir el caracter de finalizacion
	// retorna 2 valores
	text, err := reader.ReadString('\n')
	if err != nil { // Cuando todo esta bien err es nil
		fmt.Println(err)
	} else {
		fmt.Printf("Hola %s\n", text)
	}
}
