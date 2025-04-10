package main

func main() {
	// Aqui ta sendo estanciado uma função que está sendo jogado para "funcaoTest"
	funcaoTest := func(test string, testInt int) {
		println(test, testInt)
	}
	test(funcaoTest) // Chamando a funcao "test", que tem como parâmetro a função: "funcaoTest".
	// Assim executando a função que foi passada como um parâmetro
}

func test(value func(string, int)) {
	value("Ryan", 21)
	// Valores declarados na função para serem exibidos
}
