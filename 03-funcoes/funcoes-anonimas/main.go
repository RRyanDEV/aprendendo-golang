package main

func main() {
	test()
}


func test() {
	// Dentro dessa função, foi estanciado uma outra função que executa aqui dentro mesmo.
	// Tornando ela Anônima, sendo ela executavel somente aqui dentro

	func(valorString string, valorInt int){
		println(valorString, valorInt)
	}("Ryan", 21)
		// Esse parenteses no final são para indicar quais valores ela recebe nos parâmetros que foram declarados
}
