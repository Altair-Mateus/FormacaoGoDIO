package tests

import (
	"calculadora/models"
	"testing"
)

func TestDivisao(t *testing.T) {
	div := models.Divisao{}
	resEsperado := 1.0

	result, err := div.Executar(10, 5, 2)
	if err != nil {
		t.Fatalf("Erro ao realizar divisão: %v", err)
	}

	if result != resEsperado {
		t.Errorf("Resultado incorreto, Esperado: %f, Obteve: %f", resEsperado, result)
	}

}

func TestDivisaoPorZero(t *testing.T) {
	div := models.Divisao{}
	resEsperado := 10.0

	result, err := div.Executar(10, 0, 2)
	if err != nil {
		t.Fatalf("Erro ao realizar divisão: %v", err)
	}

	if result != resEsperado {
		t.Errorf("Resultado incorreto, Esperado: %f, Obteve: %f", resEsperado, result)
	}

}
