package curso1

import (
	"fmt"
)

type user2 struct {
	edad     int
	nombre   string
	apellido string
}

func (user user2) nombreCompleto() string {
	return user.nombre + " " + user.apellido
}

func (user user2) setNombre(nuevoNombre string) {
	user.nombre = nuevoNombre
}

func (user *user2) setNombre2(nuevoNombre string) {
	user.nombre = nuevoNombre
}

func (user *user2) setNombre3(nuevoNombre string) {
	(*user).nombre = nuevoNombre
}

// ExecuteGoCap14Ej1 hace tal cosa.
func ExecuteGoCap14Ej1() {
	fmt.Printf("\n\nExecuteGoCap14Ej1 <=== ****\n")
	var usuario1 user2

	usuario1.nombre = "Jaime"
	usuario1.apellido = "Rodriguez"

	fmt.Print("Usuario1: ")
	fmt.Println(usuario1)
	fmt.Print("Usuario1 nombre_completo: ")
	fmt.Println(usuario1.nombreCompleto())
	fmt.Println()

	fmt.Println("Llamo al set el cual modifica una copia y no el original...")
	usuario1.setNombre("Marcos el Chino Maidana")
	fmt.Print("Usuario1: ")
	fmt.Println(usuario1)
	fmt.Print("Usuario1 nombre_completo: ")
	fmt.Println(usuario1.nombreCompleto())
	fmt.Println()

	fmt.Println("Llamo al set el cual modifica el original...")
	usuario1.setNombre2("Marcos Maidana")
	fmt.Print("Usuario1: ")
	fmt.Println(usuario1)
	fmt.Print("Usuario1 nombre_completo: ")
	fmt.Println(usuario1.nombreCompleto())
	fmt.Println()

	fmt.Println("Llamo al OTRO set el cual modifica el original...")
	usuario1.setNombre3("Marcos el Chino Maidana")
	fmt.Print("Usuario1: ")
	fmt.Println(usuario1)
	fmt.Print("Usuario1 nombre_completo: ")
	fmt.Println(usuario1.nombreCompleto())
	fmt.Println()
}
