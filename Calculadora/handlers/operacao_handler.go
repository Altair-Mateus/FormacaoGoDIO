package handlers

import (
	"calculadora/models"
	"errors"
)

func ObterOperacao(nome string) (models.Operacao, error) {
	switch nome {
	case "soma":
		return models.Soma{}, nil

	case "subtracao":
		return models.Subtracao{}, nil

	case "multiplicacao":
		return models.Multiplicacao{}, nil

	case "divisao":
		return models.Divisao{}, nil

	default:
		return nil, errors.New("operação inválida")
	}
}
