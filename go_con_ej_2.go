package main

import (
    "fmt"
)

// tipos de variables: int, string, bool

const (
     package_priv_const_1 string = "private_const"
     Package_publ_const_2 string = "public_const"
)

func main() {
    probandoMap()
}

type elemento struct {
    value_1 string
}

func probandoMap() {
    fmt.Print("probandoMap() ======> ")
    map_1 := make(map[interface{}]interface{})
    map_1["key_1"] = "value_1"
    map_1["key_2"] = "value_2"
    fmt.Printf("map_1=%v | ", map_1)
    for key, value := range map_1 {
        fmt.Printf("\n    map_1 --> key=%s, value=%s", key, value)
    }
    fmt.Print("\n    Agregando estructuras como clave")
    map_1[elemento{"struct_val_1_1"}]="dummy_1"
    map_1[elemento{"struct_val_1_2"}]="dummy_2"
    for key, value := range map_1 {
        fmt.Printf("\n    map_1 --> key=%v, value=%v", key, value)
    }
    fmt.Println("\n    Buscando en el map.")
    fmt.Printf("    key=\"key_2\"       value=\"%v\"\n",map_1["key_2"])
    fmt.Printf("    key=struct.id     value=\"%v\"\n",map_1[elemento{"struct_val_1_1"}] )
    fmt.Printf("    key=\"struct.id\"   value=\"%v\"\n",map_1["struct_val_1_1"])
}
