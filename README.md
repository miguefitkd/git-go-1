# Pruebas de git/github y go :shipit:

## CAPITULO 1
###### INTRO
* Infraestructura dentro de las empresas...
* Lenguaje de Programacion Open Source
* Caracteristicas/Ventajas
  * Sintaxis estandarizada.
  * Compilador obliga a seguir buenas practicas.
  * Compilador muy rapido.
  * Soporte por concurrencia y facil de manejar (sintaxis muy sencilla)
  * Facil de levantar un servidor web (libreria HTTP)
  * Performance alto. Compila a lenguaje maquina. Es un poquito mas lento que C porque tiene un garbage collector
  * Se sube el binario a produccion y sale andando
* Utilidades
  * CLI TOOLS (command line tools), cosas que usamos por linea de comando que nos ayudan a automatizar un monton de cosas, infraestructura en el Backend
  * Herramientas internas de Software, Servicios WEB, tambien sirve para hacer aplicaciones web.
* GO no es un lenguaje orientado a objetos.

## CAPITULO 2 - INSTALACION Y HOLA MUNDO
* url: https://golang.org/
* Descargar el ejecutable. Instalarlo.
* Una alternativa es la consola online. Tiene ciertas limitaciones obviamente.

###### SINTAXIS
* Debe existir el "package main"
* Los imports tienen la libreria a importar entre comillas dobles "fmt" (para los print)
* Siempre debe existir un func main() {...}
* No es necesario agregar ";" al final de cada linea. El compilador las agrega.

###### COMPILAR :cyclone:
* go build [nombre_archivo].go
* Te genera un [nombre_archivo].exe
* Transforma el codigo a un binario que no es dependiente del sistema operativo.
* Se puede ejecutar desde cualquier sistema operativo.

###### EJECUTAR :running:
* Windows(consola): [nombre_archivo].exe
* Linux(consola): ./[nombre_archivo]

###### EJEMPLOS
* :computer: *[hola mundo](https://github.com/miguefitkd/git-go-1/blob/master/curso1/go_2_1.go)*

## CAPITULO 3 - VARIABLES
* Lenguaje fuertemente tipeado. El tipo de dato de una variable no cambia.
* var [nombre_variable_1, ... , nombre_variable_n] [tipo_de_dato]
* Tipado dinamico de GO, Operador ":=", la ventaja esta en que no se define el tipo de dato, se define el tipo de dato en tiempo de ejecucion en base al valor asignado.

###### Ejemplo
```go
x := 23
```
* Las variables siempre estan inicializadas,
  * int -----> 0
  * string --> cadena vacia
  * bool ----> falso
* Hay arreglos, otras cosas de slices...
* Cuando las variables no se utilizan no compila.
* Una variable no puede volver a ser definida, error de compilacion

###### CORRER DIRECTAMENTE
* go run [nombre_archivo].go

###### EJEMPLOS
* :computer: *[variables](https://github.com/miguefitkd/git-go-1/blob/master/curso1/go_3_1.go)*

## CAPITULO 4 (Conversion de tipos, casteos)
###### SINTAXIS
* Importar multiples paquetes (tambien se pueden declarar de forma individual)
```go
import(
    "fmt"
    "strconv"
)
```
* No se puede hacer un import de una libreria que no se este utilizando
* Una funcion de GO puede retornar mas de 1 valor de retorno (para descartarlo usar "_"). 
  * **Multiple values strconv.Atoi(...) in simple value contex** significa que se estan devolviendo mas valores de los que estas esperando.

###### CONVERSIONES ("strconv")
* Itoa --> int a string --> var_str = strconv.Itoa(var_int)
* Atoi --> string a int --> var_int, _ = str
* Para el fmt.Println no se pueden mezclar tipos.
  * CODIGO ---> fmt.Println("Mi edad es: " + edad)
  * ERROR ----> (type untyped string) as type int
  * SOLUCION -> edad_str := strconv.Itoa( edad )

###### EJEMPLOS
* :computer: *[conversion](https://github.com/miguefitkd/git-go-1/blob/master/curso1/go_4_1.go)*

## CAPITULO 5 (Leer e imprimir datos)
###### SINTAXIS
* Comillas simples definen un caracter
* Comillas dobles definen una cadena
* nil es NULL

###### IMPRIMIR
* fmt.Print("Hola Mundo!")
* fmt.Println("Hola Mundo!")
* fmt.Printf("Mi edad es: %d", edad)
  * %s --> imprime un string
  * %d --> decimal
  * %v --> es una especie de comodin, te retorna el valor por default, se puede usar con tranquilidad para los tipos basicos
  * %f --> Para los doubles, float, etc

###### LEER
* fmt.Scanf: Esta es la forma mas sencilla de leer agregar el salto de linea al final, pasarle el puntero de la variable. Si la entrada se define como %v se va a manejar como si fuese un %s fmt.Scanf("%d\n", &[variable_int]) el \n define el caracter de fin de lectura.
* "bufio", "os":
```go
// Construir un reader. Hay que indicarle cual es el lector que debe utilizar.
reader := bufio.NewReader(os.Stdin)
fmt.Print("Ingresar texto: ")
// Definir el caracter(definir con comillas simples) de finalizacion; retorna 2 valores
text,err := reader.ReadString('\n')
if (err != nil) { // Cuando todo esta bien err es nil
    fmt.Println(err)
} else {
    fmt.Printf("Texto ingresado:  %s\n", text)
} 
```

###### EJEMPLOS
* :computer: *[leer e imprimir 1](https://github.com/miguefitkd/git-go-1/blob/master/curso1/go_5_1.go)*
* :computer: *[leer e imprimir 2](https://github.com/miguefitkd/git-go-1/blob/master/curso1/go_5_2.go)*


## CAPITULO 6 (Condicionales)
###### SINTAXIS
* Los parentesis del if son opcionales.
* Las llaves del if son obligatorias
* Las llaves de inicio ({) siempre en la misma linea del if
* if / else if / else

###### EJEMPLOS
* :computer: *[condicionales](https://github.com/miguefitkd/git-go-1/blob/master/curso1/go_6_1.go)*

## CAPITULO 7 (Ciclos / Unico ciclo --> for)
###### SINTAXIS
* Los parentesis no son necesarios
* For ... :)
```go
for [inicializacion];[condicion del fin];[cada vez que el ciclo termina] {
    ....
} 
```
* Se puede simular un while utilizando solo la [condicion del fin]
* break / continue funcionan de igual manera

###### EJEMPLOS
* :computer: *[for](https://github.com/miguefitkd/git-go-1/blob/master/curso1/go_7_1.go)*

## CAPITULO 8 (Arreglos y Matrices)
###### SINTAXIS
* var [nombre_variable] [n][tipo]
* Se inicializa con los valores por default
```go
var arreglo [10]int
```
* Se pueden inicializar, las posiciones no inicializadas toman el valor por default
```go
arreglo2 := [10]int{1,2,3,4}
```
* Se puede imprimir todo el array --> fmt.Println( [array] )
```go
fmt.Println(arreglo2)
```
* Longitud del array len(...)
```go
fmt.Println(len(arreglo2))
```
* Elementos particulares, se acceden mediante un indice(cuidado con la longitud)
```go
arreglo2[index]
```
* Se pueden armar matrices. var [nombre_variable] [long_1][long_2][tipo]
```go
// Dos arreglos de 3 elementos...
var matrix [2][3]int
```

###### EJEMPLOS
* :computer: *[arreglos y matrices](https://github.com/miguefitkd/git-go-1/blob/master/curso1/go_8_1.go)*


## CAPITULO 9 (Slices)
###### SINTAXIS
* var [nombre_variable][][tipo_de_dato]
```go
var sliceVacio []int
sliceInicializado := []int{1,2,3,4}
```
* Estructura construida en base a un array, que se pueden redimensionar
* Los slices se pueden inicializar; cuando no estan inicializados son nil(NULL)
```go
var sliceVacio []int
if sliceVacio == nil{
    fmt.Println("El slice esta vacio")
}
```
* Estructura de un slice contiene 3 datos [ptero|longitud|capacidad]
* Se puede crear un slice en base a un array (slicing)

###### SINTAXIS
* var_slice := var_array[index_start:index_end] (el index_end es NO inclusivo)
```go
arreglo := [4] int {1,2,3,4}
slice := arreglo[:2]     // Se toma desde el inicio hasta la posicion 2 (no inclusivo)
slice := arreglo[2:]     // Se toma desde el index 2 hasta  el final del array
slice := arreglo[1:2]    // Se toma de la posicion 1 hasta la posicion 2 (no inclusivo)
// El maximo valor es la longitud del array menos 1
```
* Longitud del slice:  len(var_slice)
* Capacidad del slice: cap(var_slice)
* La capacidad en mayor o igual a la longitud

###### EJEMPLOS
* :computer: *[slices](https://github.com/miguefitkd/git-go-1/blob/master/curso1/go_9_1.go)*

## CAPITULO 10 (Make y Append)
###### SINTAXIS
* Make --> make([][tipo_del_slice],[longitud_del_slice],[capacidad_del_slice])
```go
// Crea un slice y lo inicializa.
make([]int,3) // La capacidad del slice no es requerida.
make([]int,3, 5)
make([]int,0) // slice con array de logitud cero
```
* Append --> var_slice = append(var_slice, [new_data])
```go
// Agrega un nuevo alemento junto con un valor al final del array interno.
// La longitud va a aumentar en 1.
// Cuando la longitud es menor a la capacidad solo se agrega un elemento al array.
// Cuando la logitud es igual a la capacidad se crea un nuevo slice(menor performance).
slice_int = append(slice_int, 2)
```

###### ESTRUCTURA DEL SLICE --> [ptero|longitud|capacidad]
* ptero:     Puntero al arreglo interno
* longitud:  La longitud del arreglo interno
* capacidad: Cual es la capacidad del slice

###### EJEMPLOS
* :computer: *[make and append](https://github.com/miguefitkd/git-go-1/blob/master/curso1/go_10_1.go)*

## CAPITULO 11 (Copy)
###### SINTAXIS
* var_copy = copy(destino, fuente)

* La funcion copy copia el minimo de elementos de los dos slices. El minimo len(...).
* Truco, utilizar la longitud del slice de origen de los datos
```go
slice1 := []int{1,2,3,4}
slice2 := make([]int,len(slice1))
```
* Se suele duplicar la capacidad
```go
slice1 := []int{1,2,3,4}
slice2 := make([]int,len(slice1), cap(slice1)*2)
```

###### EJEMPLOS
* :computer: *[slices - copy](https://github.com/miguefitkd/git-go-1/blob/master/curso1/go_11_1.go)*

## CAPITULO 12 (Punteros)
###### SINTAXIS
* *[tipo_de_dato]
```go
var puntero *int
```
* El valor cero (0) is nil (NULL)
* Para obtener la direccion de memoria --> &[variable]
* Acceder a la variable apuntada por el puntero --> *[puntero]

###### EJEMPLOS
* :computer: *[pointers](https://github.com/miguefitkd/git-go-1/blob/master/curso1/go_12_1.go)*

## CAPITULO 13 (Structs)
###### SINTAXIS
* Struct
```go
type nombre-tipo-struct struct {
    atributo1 int
    atributo2 string
    atributo3 string
}
var nombre-variable nombre-tipo-struct]
```
* Se inicializas los elementos del struct con los valores con default.
* Creacion.
```go
usuario1 := User{nombre : "Miguel Angel", apellido: "Cisneros"}
usuario2 := User{22, "Miguel Angel", "Da Vinci"} // especificar en orden cada dato
ptero_usuario3 := new(User)                      // retorna un puntero a la estructura
```
* Acceso a elementos de la estructura mediante el "." --> [nombre_variable].[nombre_campo_del_struct]
* Acceso a elementos de la estructura mediante punteros, tambien se puede usar la forma anterior --> (*[nombre_ptero]).[nombre_campo_del_struct]

###### EJEMPLOS
* :computer: *[structs](https://github.com/miguefitkd/git-go-1/blob/master/curso1/go_13_1.go)*

## CAPITULO 14 (Metodos* Agregando metodos a las estructuras)
###### SINTAXIS
* func ([identif_var_func] [nombre_struct]) [nombre_funcion]() [tipo_de_dato_retorno] { ... }
```go
type User struct {
    edad int
    nombre,apellido string
}
func (this User) nombre_completo() string {
    return this.nombre + " " + this.apellido
}

fmt.Println( variable_tipo_User.nombre_completo() )

```
* Se pueden agregar funciones a structs PROPIOS, no asi las de otros packages.
* Cada vez que se pasa un argumento a una funcion, este ARGUMENTO se pasa como una COPIA, es decir que SE DUPLICA el objeto y se envia. Si modifico algo de la copia no se va a reflejar en e original.
* Se puede recibir un puntero en lugar de una copia de un struct
```go
func (this *User) set_nombre(nuevoNombre string) { ... }
// Al pasarle un puntero es mas performante frente a structs grandes.
// Si se modifica la informacion de la estructura se va a estar modificando el struct original
```

###### EJEMPLOS
* :computer: *[structs and methods](https://github.com/miguefitkd/git-go-1/blob/master/curso1/go_14_1.go)*

## CAPITULO 15 (Campos anonimos para las estructuras)
###### SINTAXIS
* type [nombre_struct_1] struct { ... }
* type [nombre_struct_2] struct { [nombre_struct_1] ... }
```go
type Human struct {
    edad int
    nombre string
    apellido string
}
type Tutor struct {
    Human
}
...
tutor := Tutor{Human{"Miguel"}}
fmt.Println(tutor.nombre)
fmt.Println(tutor.Human.nombre)
...
```
* Los campos anonimos permiten replicar el comportamiento de la herencia de POO
* Para acceder se usa el "."
* Cuando hayan mas de un tipo en la estructura final puede ser que se repitan campos, asi que se puede utilizar la forma [nombre_variable].[nombre_struct_1].[campo]
```go
tutor.Human.edad
```
* Se puede agregar funciones relacionadas a las distintas estructuras e ir llamando a los metodos de la estructura que se desee utilizando el "."
* Se puede sobre escribir funciones, si existen dos funciones con identico nombre para Human y para Tutor, estaria tomando siempre la funcion de Tutor, a menos que se especifique la funcion por medio del "."
```go
func (this Tutor) hablar() string {
    return "El Tutor dice: andale la osa!!!, pero el Humano dice: " + this.Human.hablar()
}
```

###### EJEMPLOS
* :computer: *[campos anonimos](https://github.com/miguefitkd/git-go-1/blob/master/curso1/go_15_1.go)*

## CAPITULO 16 (Interfaces)
###### SINTAXIS
* Interface
  * Estructura que define metodos que no estan implementados. 
  * Es un tipo de dato.
```go
type nombre_interface  interface {
    nombre_metodo]() tipo_de_dato_devuelto
    nombre_metodo]() tipo_de_dato_devuelto
}

type User  interface {
    Permisos() int   // 1 al 5
    Nombre() string
}

type Admin struct {
    nombre string
}

func (this Admin) Permisos() int{
    return 5
}

func (this Admin) Nombre() string{
    return this.nombre
}

//Utilizando interfaz como tipo de dato, polimorfismo
func auth(user User) string{
    if user.Permisos()>4 {
        return "Este tiene permisos de ADMIN."
    }
    return "A este no le permitas hacer nada."
}
```
* Las interfaces se pueden implementar relacionados a una estructura
* No hay una palabra reservada para la implementacion de las interfaces, con el solo hecho de estar los metodos es suficiente.
* Se puede utilizar estructuras como interfaces, para obtener polimorfismo :P 
* Se puede trabajar con arreglos/slices con elementos del tipo de la interfaz (polimorfismo)

###### EJEMPLOS
* :computer: *[interfaces](https://github.com/miguefitkd/git-go-1/blob/master/curso1/go_16_1.go)*


## CAPITULO 17 (GO Routines / ejecuciones concurrentes / Como threads en java?)
###### SINTAXIS
* go [nombre_funcion([parametros_de_funcion])]
```go
go func(){
    // bloque de codigo que se ejecuta de forma concurrente mediante una funcion anonima
}()
```
* La palabra "go" mueve la ejecucion de la funcion a una GO routine. Una ejecucion separada.
* GO no utiliza un Thread del sistema operativo por funcion a ejecutar. Utiliza lo que se llama GO routines, en una computadora hay un limite de Threads del sistema que se pueden utilizar.
* Puedo tener miles de GO Routines sin que se bloquee, lo que se hace es levantar un nuevo Threaddel SO solo cuando multiples GO routines se estan bloqueando.
* Ejemplos de bloqueos: esperando la entrada de datos de un usuario, llamada a un sleep.
* Go routines son performantes. El Compilador se encarga de abrirlas y terminarlas.
* Puede causar que el programa finalice antes de terminar con todas las operaciones a ejecutar.
* Para evitar la finalizacion temprana(se coloca antes de la finalizacion del pgm):
```go
// Hack para que no termine el programa...
// El programa no termina hasta que le de enter
var wait string 
fmt.Scanln(&wait)
```

###### OTRAS COSAS
* split --> import "strings"
```go
var_array_str = strings.Split(var_str, str_para_split)
```
* range --> para iteraciones en arrays, retorna el indice y el elemento correspondiente
```go
for var_index,var_element := range(var_array) {
    // ...
}
```
* sleep --> import "time"
```go
time.Sleep(1000 + time.Millisecond)
```

###### EJEMPLOS
* :computer: *[go routines](https://github.com/miguefitkd/git-go-1/blob/master/curso1/go_17_1.go)*

## CAPITULO 18 (GO Channels)
###### SINTAXIS
* make(chan [tipo_de_dato])
* var_channel <* var_dato_enviar
* var_dato_recibido <- var_channel
```go
channel := make(chan string)
go func(channel chan string) {
    for { // buvle infinito, siempre va a estar esperando un input de datos
        var name string
        fmt.Scanln(&name)
        channel <- name
    }
}(channel)
msg := <- channel
fmt.Println(msg)
```
* Permiten comunicar go routines via un canal.
* Forma de enviar:  [canal] <- [dato]
* Forma de recibir: [var_recibiendo] = <- [canal]

###### EJEMPLOS
* :computer: *[go channels](https://github.com/miguefitkd/git-go-1/blob/master/curso1/go_18_1.go)*

## CAPITULO 19 (Leer Archivos V1 | Libreria "io/ioutil")
###### SINTAXIS
* [variable_con_informacion_en_bytes],[error] = ioutil.ReadFile([path_del_archivo])
```go
file_data, err := ioutil.ReadFile("./hola.txt")
if err == nil {
    fmt.Println(string(file_data))
}
```
* Si [error] es distinto de nil entonces se produjo un error
* Abre el archivo y carga el contenido del mismo en memoria (usar solo para archivos chicos)
* La informacion recibida esta en bytes, es necesario convertir la informacion

###### EJEMPLOS
* :computer: *[read files 1](https://github.com/miguefitkd/git-go-1/blob/master/curso1/go_19_1.go)*

## CAPITULO 20 (Leer Archivos V2 | Librerias: "bufio", "os")
###### SINTAXIS
* [variable_con_informacion_en_bytes],[error] = os.Open([path_del_archivo])
* [variable_scanner] := bufio.NewScanner([variable_con_informacion_en_bytes])
* [variable_linea_archivo_str] := [variable_scanner].Text()
* [variable_con_informacion_en_bytes].Close()
```go
archivo,err := os.Open("./conejo_pepito.txt")
if err == nil {
    scanner := bufio.NewScanner(archivo)
    for scanner.Scan() {
        linea := scanner.Text()
        fmt.Println(linea)
    }
}
```
* Es necesario abrir y cerrar el archivo

###### EJEMPLOS
* :computer: *[read files 2](https://github.com/miguefitkd/git-go-1/blob/master/curso1/go_20_1.go)*

## CAPITULO 21 (Defer)
###### SINTAXIS
* defer [nombre_func(parametros_func)]
```go
// defer func(){
//     ...
// }()
defer archivo.close()
defer func() {
    archivo.close()
    fmt.Println("El archivo fue cerrado.")
}
``` 
* Para evitar la terminacion abrupta de la ejecucion o un return antes de lo esperado.
* Agrega la ejecucion de una linea de codigo al stack de funciones que se deben ejecutar cuando la funcion actual retorna
* Sirve para asegurarse de que siempre se ejecute una linea de codigo.
* Similar al funcionamiento del finally (un poco similar)

###### EJEMPLOS
* :computer: *[defer](https://github.com/miguefitkd/git-go-1/blob/master/curso1/go_21_1.go)*

## CAPITULO 22 (Panic y Recover)
###### SINTAXIS
* panic(err_value)
* recover()

* panic(..)
  * Te muestra el error y el stack del error
  * Permite imprimir un error para poder identificar en que linea se produjo el error.
  * "defer" se ejecuta aunque haya un panic(...)
  * Cuando se tira un panic ya no continua ejecutando el flujo normal
* recover()
  * Se utiliza para detener un panic(), se puede utilizar junto con el "defer"
  * Permite recuperarte del pnic y continuar la ejecucion a partir del lugar que precises
  *Similar a un catch?

###### EJEMPLOS
* :computer: *[panic and recover](https://github.com/miguefitkd/git-go-1/blob/master/curso1/go_22_1.go)*

## CAPITULO 23 (WEB / net/http, io)
* Sintaxis 1
```go
http.HandlerFunc(
    "parte de la url",
    func (w http.ResponseWriter, r *http.Request){
        // Primer parametro: stream de datos de escritura
        // Segundo parametro: puntero a la informacion de la peticion recibida
        
        // Que se debe hacer
            
    }
)
```
* Sintaxis 2
```go
http.HandlerFunc(
     "parte de la url",
     nombre_de_la_funcion
)
...
func nombre_de_la_funcion(w http.ResponseWriter, r *http.Request){
    // Primer parametro: stream de datos de escritura
    // Segundo parametro: puntero a la informacion de la peticion recibida
    
    // Que se debe hacer
    
}
```
* http.ListenAndServe(":8080", nil)
  * Si no hay nada antes de los ":" se toma como localhost, luego viene el puerto
  * En el segundo parametro es nil para que se utilicen los manejadores definidos con http.HandlerFunc(...)

###### EJEMPLOS
* :computer: *[web 1](https://github.com/miguefitkd/git-go-1/blob/master/curso1/go_23_1.go)*

## CAPITULO 24 (WEB / Servir un archivo y un directorio)
* Sintaxis 1
```go
http.HandleFunc(
    "/",
    func(w http.ResponseWriter, r *http.Request) {
        // 
        http.ServeFile(w, r, "path_relativo_al_archivo_a_servir")
    }
)
// El path es relativo al directorio desde donde se este EJECUTANDO el programa
```
* Sintaxis 2
```go
http.HandleFunc(
    "/",
    func(w http.ResponseWriter, r *http.Request) {
        // 
        http.ServeFile(w, r, r.URL.Path[1:])
    }
)
// Esto nos permite acceder a "todos" los archivos/directorios desde 
//    donde se esta ejecutando la app
// ej:
// localhost:8080/
// localhost:8080/index.html
// localhost:8080/go_24_1.go
```
* Servir informacion a partir de un directorio para no comprometer la seguridad del sitio ni de los datos manejados en el mismo
```go
fileServer := http.FileServer(http.Dir("public"))
http.Handle("/", http.StripPrefix("/", fileServer) )
http.ListenAndServe(":8080", nil)
```

###### EJEMPLOS
* :computer: *[web 2 1](https://github.com/miguefitkd/git-go-1/blob/master/curso1/go_24_1.go)*
* :computer: *[web 2 2](https://github.com/miguefitkd/git-go-1/blob/master/curso1/go_24_2.go)*

## CAPITULO 25 (JSON / "encoding/json")
*  json.NewEncoder( un variable del tipo http.ResponseWriter).Encode( var_tipo_struct )
* Por convension de GO con respecto a JSON, los campos de un struct con nombres que comiencen con minusculas no se tienen en cuenta.
* Se puede modificar el "title" con el que se va a ver un campo de un struct, mediante la siguiente notacion...
```go
type [nombre_struct] struct {
//  ...
    NombreCampo string `json:"nombre_campo_json"`
//  ...
}
```

###### OTRAS COSAS
* En GO, cuando se utilizan arreglos de tipos, es necesario agregar una coma por elemento de la inicializacion. En otros leguajes el ultimo elemento no lleva "coma".
```go
type Curso struct {
    Title string       `json:"nombre_profesor"`
    NumeroVideos int `json:"cantidad_cursos"`
}
type Cursos []Curso
// ...
cursos := Cursos{
    Curso{"Go", 21},
    Curso{"Java", 16},
    Curso{"Phyton", 33},
}
```

###### EJEMPLOS
* :computer: *[json](https://github.com/miguefitkd/git-go-1/blob/master/curso1/go_25_1.go)*

## CAPITULO 26 (Creacion de packages LOCALES) 
###### SINTAXIS
* Crear directorio con el nombre del package a crear
* Dentro del package crear un archivo .go que tendra la funcionalidad del package(el nombre puede variar con respecto al nombre del package)
* Las librerias no se pueden ejecutar de forma independiente. No tienen una funcion main.
* Las funciones privadas no se pueden acceder desde packages externos
* La funcion init() preconfigura el package para que tenga toda la informacion que requiera
* Los atributos publicos se pueden acceder desde otros packages

###### OTRAS COSAS
* Funciones privadas --> El nombre comienza con minuscula
* Funciones publica  --> El nombre comienza con mayuscula
* Atributos privados --> El nombre comienza con minuscula
* Atributos publicos --> El nombre comienza con mayuscula

###### EJEMPLOS
* :computer: *[packages locales](https://github.com/miguefitkd/git-go-1/blob/master/curso1/go_26_1.go)*

## VARIAS PRUEBAS
###### EJEMPLOS
* :computer: *[pruebas 1](https://github.com/miguefitkd/git-go-1/blob/master/pruebas1/go_con_ej_1.go)*
* :computer: *[pruebas 2](https://github.com/miguefitkd/git-go-1/blob/master/pruebas1/go_con_ej_2.go)*

