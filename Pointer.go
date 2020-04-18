package main

import "fmt"

func main() {
	a := 42
	b := a

	a = 27
	fmt.Println(a, b)
	//Para punteros

	var aa int = 42
	var bb *int = &aa
	fmt.Println(aa, *bb)

}
