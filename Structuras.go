package main

import (
	"fmt"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
	episodes   []string
}

//En go no hay herencia, pero lo que existe es la composciono
type Animal struct {
	Name   string
	Origin string
}

//Entonces si queremos heredar de otra clase lo que podemos hace es agragar el campo detro de la estructura
//Como en este caso donde estamos agrgando Animal como caracteristica
type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {

	//LA forma de inicilizar la structura es de la siguiente manera
	aDoctor := Doctor{
		number:    3,
		actorName: "John Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith"}}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.number)
	fmt.Println(aDoctor.companions)
	fmt.Println(aDoctor.companions[1])

	//Para ver la parte de la composicion
	fmt.Println("Estes es para ver la herencia dentro de go")
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)

}
