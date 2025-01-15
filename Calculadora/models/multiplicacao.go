package models

type Multiplicacao struct{}

func (m Multiplicacao) Executar(numeros ...float64) (float64, error) {

	if err := ValidaEntrada(numeros, "multiplicação"); err != nil {
		return 0, err
	}

	mult := numeros[0]

	for _, num := range numeros[1:] {
		mult *= num
	}

	return mult, nil
}
