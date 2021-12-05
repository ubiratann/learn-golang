package main

import "testing"

func TestOla(t *testing.T) {
	resultado := OlaMundo()
	esperado := "Olá mundo"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}

func TestOlaVoce(t *testing.T) {
	resultado := OlaVoce("Ubiratan")
	esperado := "Olá Ubiratan"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
