package main

import (
	"calculadora/handlers"
	"fmt"
)

func main() {

	op, err := handlers.ObterOperacao("divisao")

	if err != nil {
		fmt.Println("Erro: ", err)
		return
	}

	result, err := op.Executar(10, 10, 15)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("Resultado: ", result)

}
