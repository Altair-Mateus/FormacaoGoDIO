package models

import "errors"

type Divisao struct{}

func (d Divisao) Executar(numeros ...float64) (float64, error) {

	if err := ValidaEntrada(numeros, "divisão"); err != nil {
		return 0, err
	}

	div := numeros[0]

	for _, num := range numeros[1:] {
		if num == 0 {
			return 0, errors.New("Não é possível dividir por zero")
		}

		div /= num
	}

	return div, nil
}
