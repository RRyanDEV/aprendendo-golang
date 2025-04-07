package main

func main() {
	if retorno := testInit(); retorno == "test" {
		print("True")
		// Inicializando a variáviel dentro do IF, ela só fica disponivel dentro da condição
	}
	
}

func testInit() string {
	return "test"
}
