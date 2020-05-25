package main

import (
  "fmt"
  "bufio"
  "os"
)

// Leer e imprimir datos

func main() {
	// Hay que indicarle que lector tiene que utilizar
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingresa tu nombre: ")
	// Definir el caracter de finalizacion
	// retorna 2 valores
	text,err := reader.ReadString('\n')
	if (err != nil) { // Cuando todo esta bien err es nil
	    fmt.Println(err)
	} else {
	    fmt.Printf("Hola %s\n", text)
	}
	
}	
