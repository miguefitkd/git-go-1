package main

import (
    "fmt"
)

type Human struct {
    edad int
    nombre string
    apellido string
}

func (this Human) hablar() string {
    return "bla bla bla"
}

type Profesional struct {
    experiencia int
    nombre string
    apellido string
	ocupacion string
}

type Tutor struct {
    Human
	Profesional
}

func (this Tutor) hablar() string {
    return "Aguante Boca!!!"
}

func (this Tutor) hablarAlMismoTiempo() string {
    return this.Human.hablar() + ".... " + this.hablar()
}

func main() {
    tutor1 := Tutor{Human{edad: 12},Profesional{}}
	
	fmt.Print("tutor1: ")
	fmt.Println(tutor1)
	fmt.Println()
	
    tutor2 := Tutor{Human{50, "Gabriel Omar", "Batistuta"}, Profesional{30, "Gabriel Omar", "Batistuta", "Pintor"}}
	fmt.Print("tutor2: ")
	fmt.Println(tutor2)
	fmt.Println()
	
	fmt.Println("Modifico la edad del tutor.")
	tutor2.edad = 51
	fmt.Print("tutor2: ")
	fmt.Println(tutor2)
	fmt.Println()
	
	fmt.Println("Modifico otra vez la edad del tutor.")
	tutor2.Human.edad = 55
	fmt.Print("tutor2: ")
	fmt.Println(tutor2)
	fmt.Println()
	
	fmt.Println("Modifico el nombre del tutor(Human).")
	tutor2.Human.nombre = "Oscar"
	tutor2.Profesional.apellido = "Ferreyra"
	fmt.Print("tutor2: ")
	fmt.Println(tutor2)
	fmt.Println()
	
	fmt.Println("Ejecutando metodos hablar().")
	fmt.Println("Metodo hablar de Tutor: " + tutor2.hablar())
	fmt.Println("Metodo hablar de Tutor.Human: " + tutor2.Human.hablar())
	fmt.Println("Metodo hablarAlMismoTiempo de Tutor:" + tutor2.hablarAlMismoTiempo())
}