package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 500; i++ {

		fmt.Println(i)
	}
	//PAra instanciar dos varibales en el mismo for

	for i, j := 0, 0; i < 30; i, j = i+1, j+2 {

		fmt.Println(i, j)
	}
	//AQUI NO HAY WHILE, PURO FOR

	//Para delcarar un while infinito

	i := 0
	for {
		fmt.Println(i)
		i++
		if i%2 == 0 {
			break
		}
	}

	//PAra usar un for similar a python podemos usar

	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println("Indice: ", k, " Value: ", v)

	}

}
