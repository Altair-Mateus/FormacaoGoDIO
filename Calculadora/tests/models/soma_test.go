package tests

import (
	"calculadora/models"
	"testing"
)

func TestSoma_Executar(t *testing.T) {
	soma := models.Soma{}
	resEsperado := 42.9

	result, err := soma.Executar(15, 6.9, 21)

	if err != nil {
		t.Fatalf("Erro ao realizar soma: %v", err)
	}

	if result != resEsperado {
		t.Errorf("Resultado incorreto, Esperado: %f, Obteve: %f", resEsperado, result)
	}
}
