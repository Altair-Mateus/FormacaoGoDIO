package main

import "fmt"

func divisivelPor3(valorInicial, valorFinal int) {
	for i := valorInicial; i <= valorFinal; i++ {
		multiplo3 := i % 3
		multiplo5 := i % 5

		if (multiplo3 == 0) && (multiplo5 == 0) {
			fmt.Println("Pin Pan")
		} else if multiplo3 == 0 {
			fmt.Println("Pin")
		} else if multiplo5 == 0 {
			fmt.Println("Pan")
		}

	}
}

func main() {
	divisivelPor3(1, 100)
}
