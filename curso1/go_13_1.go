package main

import (
    "fmt"
)

type User struct {
    edad int
    nombre string
    apellido string
}

func main() {
    var usuario1 User
	
	fmt.Print("Usuario1(valores iniciales): ")
	fmt.Println(usuario1)
	fmt.Println()
	
	fmt.Print("Usuario runtime: ")
	fmt.Println(User{nombre : "Miguel", apellido: "Ramos"})
	fmt.Println()
	
	usuario2 := User{nombre : "Miguel Angel", apellido: "Cisneros"}
	fmt.Print("Usuario2: ")
	fmt.Println(usuario2)
	fmt.Println()
	
	usuario3 := User{22, "Miguel Angel", "Da Vinci"}
	fmt.Print("Usuario3: ")
	fmt.Println(usuario3)
	fmt.Println()
	
	fmt.Println("Creando puntero a una estructura...")
	usuario4 := new(User)
	fmt.Print("*Usuario4: ")
	fmt.Println(*usuario4)
	fmt.Print("Valor contenido por Usuario4: ")
	fmt.Println(usuario4)
	fmt.Print("Asignando Nombre al struct: ")
	(*usuario4).nombre = "Conejo Pepito"
	fmt.Println(*usuario4)
	fmt.Println()
}