package tests

import (
	"calculadora/models"
	"testing"
)

func TestSubtracao(t *testing.T) {
	sub := models.Subtracao{}
	resEsperado := -10.0

	result, err := sub.Executar(10, 10, 10)
	if err != nil {
		t.Fatalf("Erro ao realizar Subtração: %v ", err)
	}

	if result != resEsperado {
		t.Errorf("Resultado incorreto, Esperado: %f, Obteve: %f", resEsperado, result)
	}

}
