package main

import (
    "fmt"
    "time"
    "strings"
)

func main(){
        fmt.Println("Ejecutando operacion de forma secuencial.")
        mi_nombre_lento("Miguel")
        fmt.Println()
	fmt.Println("Al fin termino...")
	fmt.Println()
	fmt.Println("Ejecutando operacion de forma concurrente ('go').")
	go mi_nombre_lento("Miguel")
	
	// HACK 1.-
	// Hack para que no termine el programa...
	// El programa no termina hasta que le de enter
	// var wait string 
	// fmt.Scanln(&wait)
	
	// HACK 2.-
	// Hack para que no termine el programa...
	// El programa no termina hasta que le de enter
	go func(){
		mi_nombre_lento("Angel")
	} ()
	
	go func(){
		mi_nombre_lento("Ramos")
	} ()
	
	var wait string 
	fmt.Scanln(&wait)
}

func mi_nombre_lento(nombre string){
   // Obtiene cada una de las letras
   letras := strings.Split(nombre, "")
   
   // el primer parametro es el indice
   for _,letra := range(letras) {
       time.Sleep( 1000 * time.Millisecond)
       fmt.Print(letra + " ")
   }
}
