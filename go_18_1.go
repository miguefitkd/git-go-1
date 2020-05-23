package main

import (
    "fmt"
)

func main() {
    channel := make(chan string)
	
	// Debe coincidir la declaracion con la llamada que se hace mas abajo en la linea 
	go func(channel chan string){
	    // Esta GO routine nunca se detiene esta en un ciclo infinito
	    for {
		    var nombre string
		    fmt.Scanln(&nombre)
			// Enviar informacion al canal
			channel <- nombre
		}
	}(channel)
	
	// La rutina principal se bloquea hasta que recibe algo del canal
	// 1- Recibir informacion desde el canal
	mensaje := <- channel
	fmt.Println("1-Informacion recibida desde el canal: " + mensaje)
	
	// La rutina principal se vuelve a bloquear hasta que recibe algo del canal
	// 2- Recibir informacion desde el canal
	mensaje = <- channel
	fmt.Println("2-Informacion recibida desde el canal: " + mensaje)
	
}