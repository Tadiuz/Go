package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.google.com")

	if err != nil {
		log.Fatal(err)
	}

	robots, err := ioutil.ReadAll(res.Body)
	//LA palabra defer lo que hace es ejecutarnos esa linea hasta el final de la ejecucion del main
	//DE esa forma podemos abrir recursos y cerrarlos con la palabra defer luego luego pero se cerrara hasta el final

	defer res.Body.Close()

	if err != nil {

		log.Fatal(err)
	}
	fmt.Printf("%s", robots)

}
