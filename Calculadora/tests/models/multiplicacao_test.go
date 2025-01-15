package tests

import (
	"calculadora/models"
	"testing"
)

func TestMuliplicacao(t *testing.T) {
	mult := models.Multiplicacao{}
	resEsperado := 121.00

	result, err := mult.Executar(11, 11)
	if err != nil {
		t.Fatalf("Erro ao realizar multiplicação: %v", err)
	}

	if result != resEsperado {
		t.Errorf("Resultado incorreto, Esperado: %f, Obteve: %f", resEsperado, result)
	}

}
