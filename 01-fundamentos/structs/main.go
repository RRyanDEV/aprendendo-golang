package main

import "fmt"

func main() {
	var user User = User{
	// Forma de declarar valores de um objeto.
		name: "Ryan",
		age: 21,
	}

	fmt.Print(user)
}

type User struct{
	// Forma de instanciar um objeto, e definir suas propriedades.
	name string
	age int
}
