package main

import "fmt"

func main() {

	texto := "Olá Mundo"      //String
	numero := 21              // Number
	numeroFlutuante := 21.521 // Float
	booleean := true          //Boolean

	fmt.Printf("%T", texto)
	fmt.Println("-----")
	fmt.Printf("%T", numero)
	fmt.Println("-----")
	fmt.Printf("%T", numeroFlutuante)
	fmt.Println("-----")
	fmt.Printf("%T", booleean)
	fmt.Println("-----")
}
