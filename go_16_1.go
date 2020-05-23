package main

import (
    "fmt"
)


type User  interface {
    Permisos() int
	Nombre() string
}

type Admin struct {
    nombre string
}

func (this Admin) Permisos() int {
    return 5
}

func (this Admin) Nombre() string {
    return this.nombre
}

// --------------------------------

type Editor struct {
    nombre string
}

func (this Editor) Permisos() int {
    return 3
}

func (this Editor) Nombre() string {
    return this.nombre
}

// --------------------------------

func auth(usuario User) string {
    if (usuario.Permisos() >= 5) {
	   return usuario.Nombre() + " tiene permisos de Administrador"
	} else {
	   return usuario.Nombre() + " NO tiene permisos de Administrador"
	}
}


func main(){
    admin := Admin{"Miguel Angel Ramos"}
	fmt.Println("Autenticando usuario: " + admin.nombre)
	// La estructura Admin esta implementando la interfaz
	fmt.Println(auth(admin))
	fmt.Println()
	
	
	editor := Editor{"Micaela Angelica Ramos"}
	fmt.Println("Autenticando usuario: " + editor.nombre)
	// La estructura Editor esta implementando la interfaz
	fmt.Println(auth(editor))
	fmt.Println()
	
	
	fmt.Println("Creando e iterando un Slice de Users...")
	usuarios := [] User{Admin{"Celmira del Carmen Cisneros"},Admin{"Sebastian Baudilio Ramos"},Editor{"Federico Alejandro Ramos"}}
	for _,usuario := range usuarios {
	    fmt.Println(auth(usuario))
	}
	fmt.Println()
}