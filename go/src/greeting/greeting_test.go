package main

import "testing"

func TestGreeting(t *testing.T)  {

	got := greeting("teste")
	want := "<b>teste</b>"

	if got != want {
		t.Errorf("funcao Greeting não esta certa.\n devolveu: %v \n esperado: %v", got, want)
	}
	
}