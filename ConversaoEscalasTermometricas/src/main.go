package main

import "fmt"

const PONTOEBULICAOF float64 = 212.0

func main() {

	PontoEbulicaoC := ((PONTOEBULICAOF - 32) / 1.8)

	fmt.Printf("Ponto de Ebulição em Fahrenheit: %.2f \nPonto Ebulição em Celsius: %.2f\n", PONTOEBULICAOF, PontoEbulicaoC)

}
