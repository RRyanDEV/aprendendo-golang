package main

func main() {
	valoress("Ryan", "Miguel", "Dex", "Mestre", "Ice", "Lucas", "Rodrigo")
}

func valoress(valoresString ...string) {
	// (0 à N) valores de string, se tornando opção a quantidade de valores que vai ser passado.

	for _, x := range valoresString {
		println(x)
	}
}
