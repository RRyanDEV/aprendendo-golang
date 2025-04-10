// Ele vai executar não importa o que aconteça na primeira execução
// Vai fazer uma verificação até cair na condição declarada

package main

func main() {
	anExpression := false

	for ok := true; /*(Estanciar a variável que seria validada)*/ ok; /*(Aqui faz com que inicie o loop)*/ ok = anExpression /*(Aqui entra a condição)*/ {
		println("Passou aqui")
	}
}
