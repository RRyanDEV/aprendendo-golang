package main

import "fmt"

func main() {
	// Array é uma lista com tamanho definido.
	var arr [5]string = [5]string{"test", "test2", "test3", "test4", "test5"}

	// Slice é uma lista sem tamanho definido. Conforme adiciona valores, o tamanho muda.
	var slice []string = []string{"testSlice", "testSlice2", "testSlice3", "testSlice4", "testSlice5"}

	fmt.Println(arr, "- Valores dentro do Array")
	fmt.Println(len(arr), "- Tamanho do Array")
	fmt.Println(cap(arr), "- Tamanho que o Array suporta")
	fmt.Println("---------")
	/* arr = append(arr, "RyanTest") / slice = append(slice, "Ryan")
	1 . Metodo para adicionar um valor dentro de um Array ou Slice.
	2. Ele obedece o parâmetro na hora de declarar. Então caso o array ja esteja completo ele não adiciona valores.
	3. Ao adicionar um valor no Slice ele aloca um tamanho fixo e ao adicionar um valor ele dobra de tamanho.
	*/
	fmt.Println("---------")
	fmt.Println(slice, "- Valores dentro do Slice")
	fmt.Println(len(slice), "- Tamanho que o Slice tem")
	fmt.Println(cap(slice), "- Tamanho que o Slice Suporta/Aloca.")
	fmt.Println("---------")
	slice = append(slice, "Ryan")
	fmt.Println(slice, "- Valores dentro do Slice, depois de acresentar um valor")
	fmt.Println(len(slice), "- Tamanho que o Slice tem, depois de acresentar um valor")
	fmt.Println(cap(slice), "- Tamanho que o Slice Suporta/Aloca, depois de acresentar um valor")
	fmt.Println("---------")
}
