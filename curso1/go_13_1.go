package curso1

import (
	"fmt"
)

type user1 struct {
	edad     int
	nombre   string
	apellido string
}

// ExecuteGoCap13Ej1 hace tal cosa.
func ExecuteGoCap13Ej1() {
	fmt.Printf("\n\nExecuteGoCap13Ej1 <=== ****\n")
	var usuario1 user1

	fmt.Print("Usuario1(valores iniciales): ")
	fmt.Println(usuario1)
	fmt.Println()

	fmt.Print("Usuario runtime: ")
	fmt.Println(user1{nombre: "Miguel", apellido: "Ramos"})
	fmt.Println()

	usuario2 := user1{nombre: "Miguel Angel", apellido: "Cisneros"}
	fmt.Print("Usuario2: ")
	fmt.Println(usuario2)
	fmt.Println()

	usuario3 := user1{22, "Miguel Angel", "Da Vinci"}
	fmt.Print("Usuario3: ")
	fmt.Println(usuario3)
	fmt.Println()

	fmt.Println("Creando puntero a una estructura...")
	usuario4 := new(user1)
	fmt.Print("*Usuario4: ")
	fmt.Println(*usuario4)
	fmt.Print("Valor contenido por Usuario4: ")
	fmt.Println(usuario4)
	fmt.Print("Asignando Nombre al struct: ")
	(*usuario4).nombre = "Conejo Pepito"
	fmt.Println(*usuario4)
	fmt.Println()
}
