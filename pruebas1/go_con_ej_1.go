package pruebas1

import (
	"fmt"
	"time"
)

// tipos de variables: int, string, bool

const (
	packagePrivConst1 string = "private_const"
	packagePublConst2 string = "public_const"
)

// ExecuteGoConEj1 hace tal cosa
func ExecuteGoConEj1() {
	fmt.Printf("\n\nExecuteGoConEj1 <=== ****\n")
	probandoConstantes()
	probandoFor(3)
	probandoIf(5)
	probandoSwitch(5)
	probandoArray()
	probandoSlice()
	probandoMap()
	probandoRange()
}

func probandoConstantes() {
	fmt.Print("probandoConstantes() ======> ")
	fmt.Print("packagePrivConst1=" + packagePrivConst1 + " | ")
	fmt.Println("packagePublConst2=" + packagePublConst2 + " | ")
}

func probandoFor(parameter int) {
	fmt.Print("probandoFor() ======>")

	var i int = parameter
	for i <= 3 {
		fmt.Printf(" index_i=%d |", i)
		i++
	}

	for j := 1; j <= parameter; j++ {
		fmt.Printf(" index_j=%d |", j)
	}

	var f int = parameter
	for {
		fmt.Printf(" index_f=%d |", f)

		// Este if tiene 2 instrucciones
		if f > 0 {
			f--
		} else {
			fmt.Println("It is time to break")
			break
		}
	}
}

func probandoIf(parametro int) {
	fmt.Printf("probandoIf() ======> parametro=%d | ", parametro)
	if index := 9; parametro < index {
		fmt.Println("El parametro es menor a 9")
	} else if parametro > index {
		fmt.Println("El parametro es mayor a 9")
	} else {
		fmt.Println("El parametro es igual a 9")
	}
}

func probandoSwitch(parametro int) {
	fmt.Printf("probandoSwitch() ======> parametro=%d | ", parametro)

	switch parametro {
	case 1:
		fmt.Print("case ")
		fmt.Print("1: el parametro es 1 | ")
	case 2:
		fmt.Print("case 2: el parametro es 2 | ")
	case 3:
		fmt.Print("case 3")
		fmt.Print(":")
		fmt.Print(" el parametro es 3 |")
	default:
		fmt.Print("default: el valor no esta contemplado | ")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Print("It is ")
		fmt.Printf("Weekend (%v)| ", time.Now().Weekday())
	default:
		fmt.Printf("It is Week day (%v) | ", time.Now().Weekday())
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Printf("It is before noon (%v) |\n", t.Hour())
	default:
		fmt.Printf("It is after noon (%v) |\n", t.Hour())
	}
}

func probandoArray() {
	fmt.Print("probandoArray() ======> ")

	var array1 [5]int
	array1[2] = 100

	array2 := [4]int{1, 5, 9}

	fmt.Printf("array1: %v | array2: %v |\n", array1, array2)
}

func probandoSlice() {
	fmt.Print("probandoSlice() ======> ")
	slice1 := make([]int, 3)
	fmt.Printf("slice1=%v | ", slice1)
	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 3
	fmt.Printf("slice1=%v | \n", slice1)
	fmt.Printf("         len(slice1)=%d | ", len(slice1))
	fmt.Printf("cap(slice1)=%d | \n", cap(slice1))
	slice1 = append(slice1, 10)
	slice1 = append(slice1, 10, 20, 30)
	fmt.Printf("         slice1=%v | ", slice1)
	// Creando slice a medida
	// Copiando de slice1 al slice2
	// copy copia la minima cantidad posible
	slice2 := make([]int, len(slice1), cap(slice1))
	copy(slice2, slice1)
	fmt.Printf("slice2=%v | \n", slice2)
	// Creando slices a partir de otro slice
	slice3 := slice1[1:3]
	fmt.Printf("         slice3=%v | ", slice3)
	slice4 := slice1[:3]
	fmt.Printf("slice4=%v | \n", slice4)
	slice5 := slice1[2:]
	fmt.Printf("         slice5=%v | ", slice5)
	// Slice inicializado
	slice6 := []int{1, 2, 4, 6}
	fmt.Printf("slice6=%v | \n", slice6)
}

func probandoMap() {
	fmt.Print("probandoMap() ======> ")
	map1 := make(map[string]string)
	map1["key1"] = "value1"
	map1["key2"] = "value2"
	fmt.Printf("map1=%v | \n", map1)
	map1["key2"] = "value2_2"
	fmt.Print("         Value to key1=" + map1["key1"] + " | \n")
	fmt.Print("         Value to key_n=" + map1["key_n"] + " | \n")
	delete(map1, "key2")
	fmt.Printf("         map1=%v | \n", map1)
	value1, value1Present := map1["key1"]
	value2, value2Present := map1["key2"]
	fmt.Printf("         Value to key1=%v | ", value1)
	fmt.Printf("Value to prs_1=%v | \n", value1Present)
	fmt.Printf("         Value to key_n=%v | ", value2)
	fmt.Printf("Value to prs_1=%v | \n", value2Present)
}

func probandoRange() {
	fmt.Print("probandoRange() ======> ")
	array1 := []string{"Hola ", "Mundo!", "co te va!"}
	slice1 := []string{"Mi ", "nombre ", "es", "Miguel"}
	map1 := make(map[string]string)
	map1["key1"] = "value1"
	map1["key2"] = "value2"
	for index, value := range array1 {
		fmt.Printf("\n    array1 --> index=%d, value=%v", index, value)
	}
	for index, value := range slice1 {
		fmt.Printf("\n    slice1 --> index=%d, value=%v", index, value)
	}
	for key, value := range map1 {
		fmt.Printf("\n    map1 --> key=%s, value=%s", key, value)
	}
	fmt.Println("")
}
