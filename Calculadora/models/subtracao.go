package models

type Subtracao struct{}

func (s Subtracao) Executar(numeros ...float64) (float64, error) {

	if err := ValidaEntrada(numeros, "subtração"); err != nil {
		return 0, err
	}

	sub := numeros[0]

	for _, num := range numeros[1:] {
		sub -= num
	}

	return sub, nil
}
