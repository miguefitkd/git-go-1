package main

import (
    "fmt"
)

func main() {
   var arreglo1 [10]int
   for i:=0;i<10;i++ {
       fmt.Printf("Valor[%d] : %d\n", i, arreglo1[i])
   }
   fmt.Println();
   
   // 
   arreglo2 := [10]int{1,2,3,4}
   for i:=0;i<10;i++ {
       fmt.Printf("Valor[%d] : %d\n", i, arreglo2[i])
   }
   fmt.Println();
   
   fmt.Println(arreglo2)
   fmt.Println();
   
   fmt.Print("Longitud del array: ")
   fmt.Println(len(arreglo2))
   fmt.Println();
   
   for i:=0;i<len(arreglo2);i++ {
       fmt.Printf("Valor[%d] : %d\n", i, arreglo2[i])
   }
   fmt.Println();
   // -----------------
   fmt.Println("--- MATRIX ---");
   var matrix1 [2][3] int
   fmt.Println(matrix1);
   
   matrix1 [0][0] = 5
   matrix1 [0][1] = 7
   matrix1 [0][2] = 8
   fmt.Println(matrix1);
   
}