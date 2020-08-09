package curso1

import (
	"fmt"
	"strings"
	"time"
)

// ExecuteGoCap17Ej1 hace tal cosa.
func ExecuteGoCap17Ej1() {
	fmt.Printf("\n\nExecuteGoCap17Ej1 <=== ****\n")
	fmt.Println("Ejecutando operacion de forma secuencial.")
	miNombreLento("Miguel")
	fmt.Println()
	fmt.Println("Al fin termino...")
	fmt.Println()
	fmt.Println("Ejecutando operacion de forma concurrente ('go').")
	go miNombreLento("Miguel")

	// HACK 1.-
	// Hack para que no termine el programa...
	// El programa no termina hasta que le de enter
	// var wait string
	// fmt.Scanln(&wait)

	// HACK 2.-
	// Hack para que no termine el programa...
	// El programa no termina hasta que le de enter
	go func() {
		miNombreLento("Angel")
	}()

	go func() {
		miNombreLento("Messi")
	}()

	var wait string
	fmt.Scanln(&wait)
}

func miNombreLento(nombre string) {
	// Obtiene cada una de las letras
	letras := strings.Split(nombre, "")

	// el primer parametro es el indice
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Print(letra + " ")
	}
}
