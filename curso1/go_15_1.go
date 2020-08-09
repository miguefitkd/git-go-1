package curso1

import (
	"fmt"
)

type human struct {
	edad     int
	nombre   string
	apellido string
}

func (human human) hablar() string {
	return "bla bla bla"
}

type profesional struct {
	experiencia int
	nombre      string
	apellido    string
	ocupacion   string
}

type tutor struct {
	human
	profesional
}

func (tutor tutor) hablar() string {
	return "Aguante Boca!!!"
}

func (tutor tutor) hablarAlMismoTiempo() string {
	return tutor.human.hablar() + ".... " + tutor.hablar()
}

// ExecuteGoCap15Ej1 hace tal cosa.
func ExecuteGoCap15Ej1() {
	fmt.Printf("\n\nExecuteGoCap15Ej1 <=== ****\n")
	tutor1 := tutor{human{edad: 12}, profesional{}}

	fmt.Print("tutor1: ")
	fmt.Println(tutor1)
	fmt.Println()

	tutor2 := tutor{human{50, "Gabriel Omar", "Batistuta"}, profesional{30, "Gabriel Omar", "Batistuta", "Pintor"}}
	fmt.Print("tutor2: ")
	fmt.Println(tutor2)
	fmt.Println()

	fmt.Println("Modifico la edad del tutor.")
	tutor2.edad = 51
	fmt.Print("tutor2: ")
	fmt.Println(tutor2)
	fmt.Println()

	fmt.Println("Modifico otra vez la edad del tutor.")
	tutor2.human.edad = 55
	fmt.Print("tutor2: ")
	fmt.Println(tutor2)
	fmt.Println()

	fmt.Println("Modifico el nombre del tutor(Human).")
	tutor2.human.nombre = "Oscar"
	tutor2.profesional.apellido = "Ferreyra"
	fmt.Print("tutor2: ")
	fmt.Println(tutor2)
	fmt.Println()

	fmt.Println("Ejecutando metodos hablar().")
	fmt.Println("Metodo hablar de Tutor: " + tutor2.hablar())
	fmt.Println("Metodo hablar de Tutor.Human: " + tutor2.human.hablar())
	fmt.Println("Metodo hablarAlMismoTiempo de Tutor:" + tutor2.hablarAlMismoTiempo())
}
