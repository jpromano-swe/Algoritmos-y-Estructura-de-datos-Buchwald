package main

import "tdas/pila"

func procesarOperacion(token string, pilaOperadores pila.Pila[string], salida *[]string) {
	for !pilaOperadores.EstaVacia() && debeDesapilarOperador(pilaOperadores.VerTope(), token) {
		*salida = append(*salida, pilaOperadores.Desapilar())
	}
	pilaOperadores.Apilar(token)
}

func procesarParentesisDeCierre(token string, pilaOperadores pila.Pila[string], salida *[]string) {
	for !pilaOperadores.EstaVacia() && pilaOperadores.VerTope() != "(" {
		*salida = append(*salida, pilaOperadores.Desapilar())
	}
	pilaOperadores.Desapilar()
}

func vaciarPilaOperadores(pilaOperadores pila.Pila[string], salida *[]string) {
	for !pilaOperadores.EstaVacia() {
		*salida = append(*salida, pilaOperadores.Desapilar())
	}
}
