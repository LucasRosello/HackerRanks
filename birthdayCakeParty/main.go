package main

func calcularVelasASoplar(velas []int) int {
	velaMasAlta := 0
	velasSopladas := 0

	for _, v := range velas {
		if v > velaMasAlta {
			velasSopladas = 1
			velaMasAlta = v
		} else if v == velaMasAlta {
			velasSopladas++
		}
	}

	return velasSopladas
}
