package main

import "testing"

func TestSoma(t *testing.T) {

	total := Soma(20, 5)

	if total != 30 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", total, 30)
	}
}
