package main

import (
    "fmt"
)

type User struct {
    edad int
    nombre string
    apellido string
}

func (this User) nombre_completo() string {
    return this.nombre + " " + this.apellido
}

func (this User) set_nombre(nuevoNombre string) {
    this.nombre = nuevoNombre
}

func (this *User) set_nombre_2(nuevoNombre string) {
    this.nombre = nuevoNombre
}


func (this *User) set_nombre_3(nuevoNombre string) {
    (*this).nombre = nuevoNombre
}

func main() {
    var usuario1 User
	
	usuario1.nombre = "Jaime"
	usuario1.apellido = "Rodriguez"
	
	fmt.Print("Usuario1: ")
	fmt.Println(usuario1)
	fmt.Print("Usuario1 nombre_completo: ")
	fmt.Println(usuario1.nombre_completo())
	fmt.Println()
	
	fmt.Println("Llamo al set el cual modifica una copia y no el original...")
	usuario1.set_nombre("Marcos el Chino Maidana")
	fmt.Print("Usuario1: ")
	fmt.Println(usuario1)
	fmt.Print("Usuario1 nombre_completo: ")
	fmt.Println(usuario1.nombre_completo())
	fmt.Println()
	
	fmt.Println("Llamo al set el cual modifica el original...")
	usuario1.set_nombre_2("Marcos Maidana")
	fmt.Print("Usuario1: ")
	fmt.Println(usuario1)
	fmt.Print("Usuario1 nombre_completo: ")
	fmt.Println(usuario1.nombre_completo())
	fmt.Println()
	
	fmt.Println("Llamo al OTRO set el cual modifica el original...")
	usuario1.set_nombre_3("Marcos el Chino Maidana")
	fmt.Print("Usuario1: ")
	fmt.Println(usuario1)
	fmt.Print("Usuario1 nombre_completo: ")
	fmt.Println(usuario1.nombre_completo())
	fmt.Println()
}