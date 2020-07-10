package main

import (
    "fmt"
    "time"
)

// tipos de variables: int, string, bool

const (
     package_priv_const_1 string = "private_const"
     Package_publ_const_2 string = "public_const"
)

func main() {
    probandoConstantes()
    probandoFor( 3 )
    probandoIf(5)
    probandoSwitch(5)
    probandoArray()
    probandoSlice()
    probandoMap()
    probandoRange()
}

func probandoConstantes() {
    fmt.Print("probandoConstantes() ======> ")
    fmt.Print("package_priv_const_1=" + package_priv_const_1 + " | ")
    fmt.Println("Package_publ_const_2=" + Package_publ_const_2 + " | ")
}

func probandoFor(parameter int) {
    fmt.Print("probandoFor() ======>")

    var i int = parameter
    for i<=3 {
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
            fmt.Print("It is ");
            fmt.Printf("Weekend (%v)| ", time.Now().Weekday());
        default:
            fmt.Printf("It is Week day (%v) | ", time.Now().Weekday());
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

    var array_1 [5]int
    array_1[2] = 100

    array_2 := [4]int {1, 5, 9}

    fmt.Printf("Array_1: %v | Array_2: %v |\n", array_1, array_2)
}

func probandoSlice() {
    fmt.Print("probandoSlice() ======> ")
    slice_1 := make([]int, 3)
    fmt.Printf("slice_1=%v | ", slice_1)
    slice_1[0] = 1
    slice_1[1] = 2
    slice_1[2] = 3
    fmt.Printf("slice_1=%v | \n", slice_1)
    fmt.Printf("         len(slice_1)=%d | ", len(slice_1))
    fmt.Printf("cap(slice_1)=%d | \n", cap(slice_1))
    slice_1 = append(slice_1, 10)
    slice_1 = append(slice_1, 10, 20, 30)
    fmt.Printf("         slice_1=%v | ", slice_1)
    // Creando slice a medida
    // Copiando de slice_1 al slice_2
    // copy copia la minima cantidad posible
    slice_2 := make([]int, len(slice_1), cap(slice_1))
    copy(slice_2, slice_1)
    fmt.Printf("slice_2=%v | \n", slice_2)
    // Creando slices a partir de otro slice
    slice_3 := slice_1[1:3]
    fmt.Printf("         slice_3=%v | ", slice_3)
    slice_4 := slice_1[:3]
    fmt.Printf("slice_4=%v | \n", slice_4)
    slice_5 := slice_1[2:]
    fmt.Printf("         slice_5=%v | ", slice_5)
    // Slice inicializado
    slice_6 := []int {1, 2, 4, 6}
    fmt.Printf("slice_6=%v | \n", slice_6)
}

func probandoMap() {
    fmt.Print("probandoMap() ======> ")
    map_1 := make(map[string]string)
    map_1["key_1"] = "value_1"
    map_1["key_2"] = "value_2"
    fmt.Printf("map_1=%v | \n", map_1)
    map_1["key_2"] = "value_2_2"
    fmt.Print("         Value to key_1=" + map_1["key_1"] + " | \n")
    fmt.Print("         Value to key_n=" + map_1["key_n"] + " | \n")
    delete(map_1, "key_2")
    fmt.Printf("         map_1=%v | \n", map_1)
    value_1, value_1_present := map_1["key_1"]
    value_2, value_2_present := map_1["key_2"]
    fmt.Printf("         Value to key_1=%v | ", value_1)
    fmt.Printf("Value to prs_1=%v | \n", value_1_present)
    fmt.Printf("         Value to key_n=%v | ", value_2)
    fmt.Printf("Value to prs_1=%v | \n", value_2_present)
}

func probandoRange() {
    fmt.Print("probandoRange() ======> ")
    array_1 := []string {"Hola ", "Mundo!", "co te va!"}
    slice_1 := []string {"Mi ", "nombre ", "es", "Miguel"}
    map_1   := make(map[string]string)
    map_1["key_1"] = "value_1"
    map_1["key_2"] = "value_2"
    for index, value := range array_1 {
        fmt.Printf("\n    array_1 --> index=%d, value=%v", index, value)
    }
    for index, value := range slice_1 {
        fmt.Printf("\n    slice_1 --> index=%d, value=%v", index, value)
    }
    for key, value := range map_1 {
        fmt.Printf("\n    map_1 --> key=%s, value=%s", key, value)
    }
    fmt.Println("")
}
