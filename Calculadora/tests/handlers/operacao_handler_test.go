package tests

import (
	"calculadora/handlers"
	"calculadora/models"
	"testing"
)

func TestOperacaoHandler(t *testing.T) {
	resEsperado := "divisao"
	op, err := handlers.ObterOperacao("divisao")
	if err != nil {
		t.Fatalf("Erro inesperado: %v", err)
	}

	if _, ok := op.(models.Divisao); !ok {
		t.Errorf("Esperado %s, mas obteve: %T", resEsperado, op)
	}
}
