package main

import (
	"fmt"
)

func main() {
	//UN mapa es practicamente un dicicionario

	statePopulations := map[string]int{
		"California": 123123213,
		"Texas":      1212322423,
		"Mexico":     4354354344534,
		"Illinois":   546546,
	}

	fmt.Println(statePopulations)

	//SI queremos inicializar un map pero no darle valor lo podemos hace con la funcion make

	statePopulations_2 := make(map[string]int)
	statePopulations_2 = map[string]int{
		"California": 123123213,
		"Texas":      1212322423,
		"Mexico":     4354354344534,
		"Illinois":   546546,
	}

	fmt.Println("Inicializando primero")
	fmt.Println(statePopulations_2)

	//Para borrar unn elemento

	delete(statePopulations_2, "Texas")
	fmt.Println("DEpues de eliminar TExas")
	fmt.Println(statePopulations_2)
}
