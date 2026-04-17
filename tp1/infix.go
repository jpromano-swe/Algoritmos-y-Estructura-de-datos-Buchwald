package main

import "tdas/pila"

func InfijaAPosfija(tokens []string) []string {
	var salida []string
	pilaOperadores := pila.CrearPilaDinamica[string]()

	for _, token := range tokens {
		if esOperacion(token) {
			procesarOperacion(token, pilaOperadores, &salida)
		} else if token == "(" {
			pilaOperadores.Apilar(token)
		} else if token == ")" {
			procesarParentesisDeCierre(token, pilaOperadores, &salida)
		} else {
			salida = append(salida, token)
		}
	}
	vaciarPilaOperadores(pilaOperadores, &salida)
	return salida
}
