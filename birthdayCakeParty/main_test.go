package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSinVelas(t *testing.T) {
	velasArr := []int{}

	resultado := calcularVelasASoplar(velasArr)
	resultadoEsperado := 0

	assert.Equal(t, resultado, resultadoEsperado, "Deberian ser iguales")
}

func TestConDosVelasASoplar(t *testing.T) {
	velasArr := []int{3, 3, 0, 2}

	resultado := calcularVelasASoplar(velasArr)
	resultadoEsperado := 2

	if resultado != resultadoEsperado {
		t.Errorf("Funcion calcularVelasASoplar() arrojo el resultado = %v, pero el esperado es  %v", resultado, resultadoEsperado)
	}
}

func TestConMuchasVelas(t *testing.T) {
	velasArr := []int{7, 8, 2, 3, 1, 23, 0, 1, 2, 3, 4, 5, 144, 3, 0, 2}

	resultado := calcularVelasASoplar(velasArr)
	resultadoEsperado := 1

	if resultado != resultadoEsperado {
		t.Errorf("Funcion calcularVelasASoplar() arrojo el resultado = %v, pero el esperado es  %v", resultado, resultadoEsperado)
	}
}
