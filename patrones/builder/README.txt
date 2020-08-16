Permite separar la construccion de un objeto (complejo o no) de su representacion con el fin 
de que el mismo proceso de creacion pueda crear diferentes representaciones.

Product: objeto principal construido. Representa el objeto bajo construccion.
Builder: define la interface con los metodos que deberan cumplir los constructores
Concrete Builder: implementaciones de la interface Builder para entregar el objeto concreto
Director: Construye el objeto utilizando la interface

Se debe construir el objeto y representarlo...

