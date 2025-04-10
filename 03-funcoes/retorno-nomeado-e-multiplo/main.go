package main

import "errors"

func main() {
	value, err := test()

	if err != nil {
		println(err)
	}
	println(value)
	print(retornoNomeado())
}

func test() (string, error) {
	return "Passou", errors.New("Test")

}

func retornoNomeado() (retornoString string, retornoError error) {
	retornoError = nil
	retornoString = "Ryan"
	return
}
