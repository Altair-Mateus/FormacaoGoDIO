package models

type Soma struct{}

func (s Soma) Executar(numeros ...float64) (float64, error) {

	if err := ValidaEntrada(numeros, "soma"); err != nil {
		return 0, err
	}

	soma := 0.0

	for _, num := range numeros {
		soma += num
	}

	return soma, nil

}
