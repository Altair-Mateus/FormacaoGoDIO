package models

import "errors"

type Operacao interface {
	Executar(numeros ...float64) (float64, error)
}

func ValidaEntrada(numeros []float64, operacao string) error {
	if len(numeros) == 0 {
		return errors.New("Nenhum número fornecido para " + operacao)
	}

	return nil
}
