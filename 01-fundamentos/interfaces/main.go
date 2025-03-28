package main

import "fmt"

func main(){
	// Interfaces são tipo "genérico", podendo receber qualquer valor.
	var teste interface{}

	teste = 21.5

	fmt.Printf("%T", teste)
}