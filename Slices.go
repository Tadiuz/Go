package main

import (
	"fmt"
)

func main() {

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //This is an adress
	b := a[:]
	c := a[3:]
	d := a[:6]
	e := a[3:6]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	//SLices works as lists in python
	aa := []int{}
	fmt.Println(aa)
	fmt.Printf("Lenght: %v\n", len(aa))
	fmt.Printf("Capacity: %v\n", cap(aa))
	aa = append(aa, 1)
	fmt.Println(aa)
	fmt.Printf("Lenght: %v\n", len(aa))
	fmt.Printf("Capacity: %v\n", cap(aa))

	bb := []int{}
	bb = append(bb, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(bb)
	fmt.Printf("Lenght: %v\n", len(bb))
	fmt.Printf("Capacity: %v\n", cap(bb))

	//HOw to remove elements from an slice
	//TO delete teh frist or last element

	cc := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(cc)
	dd := cc[1:]
	fmt.Println(dd)
	dd = cc[:len(cc)-1]
	fmt.Println(dd)

	//PAra eliminar un valor en especifico podems usar lo siguiente
	//EN este ejemplo vamos a remover el numero 3
	fmt.Println("___PAra eliminar valores___")
	aaa := []int{1, 2, 3, 4, 5}
	fmt.Println(aaa)
	bbb := append(aaa[:2], aaa[3:]...)
	fmt.Println(bbb)
	fmt.Println(aaa)
}
