package main

func main() {
	test := "test"

	switch test {
	case "test":
		println("Caiu na primeira condição")
		fallthrough // Indenpendente do resultado, ele vai cair na próxima condição
	case "test2", "test3":
		println("Caiu na segunda condição")
		fallthrough // Sempre se repete, caso queira que alguma condição, mesmo sendo TRUE, continue executando o código.
	case "pareAqui":
		println("Ultimo a ser impresso")
	case "treta":
		println("Não vai ser impresso")
		break // Já vem como "DEFAULT" em cada case dentro de um switch.
	}

	println("=================================")

	defaultt := "Ryan"
	switch defaultt {
	case "Hanzo":
		println(".Caiu na primeira condição")
	case "Lucas":
		println(".Caiu na segunda condição")
	default:
		println("Caiu no Default")
	}
}
