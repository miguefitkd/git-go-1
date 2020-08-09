package pruebas1

import (
	"fmt"
)

// tipos de variables: int, string, bool

const (
	packagePrivConst3 string = "private_const"
	packagePrivConst4 string = "public_const"
)

// ExecuteGoConEj2 hace tal cosa.
func ExecuteGoConEj2() {
	fmt.Printf("\n\nExecuteGoConEj2 <=== ****\n")
	probandoMap2()
}

type elemento struct {
	value1 string
}

func probandoMap2() {
	fmt.Print("probandoMap() ======> ")
	map1 := make(map[interface{}]interface{})
	map1["key1"] = "value1"
	map1["key2"] = "value2"
	fmt.Printf("map1=%v | ", map1)
	for key, value := range map1 {
		fmt.Printf("\n    map1 --> key=%s, value=%s", key, value)
	}
	fmt.Print("\n    Agregando estructuras como clave")
	map1[elemento{"structVal11"}] = "dummy_1"
	map1[elemento{"structVal12"}] = "dummy_2"
	for key, value := range map1 {
		fmt.Printf("\n    map1 --> key=%v, value=%v", key, value)
	}
	fmt.Println("\n    Buscando en el map.")
	fmt.Printf("    key=\"key2\"       value=\"%v\"\n", map1["key2"])
	fmt.Printf("    key=\"struct.id\"  value=\"%v\"\n", map1[elemento{"structVal11"}])
	fmt.Printf("    key=\"struct.id\"  value=\"%v\"\n", map1["structVal11"])
}
