package main

func main() {
	test := []string{"text1", "text2", "text3"}

	// Ele me entrega duas variáveis: um índice e a variável em sí.
	// underline(_), ignora um dos valores do retorno que vai vir.
	// 
	for _, value := range test {
		println(value)
	}
}

