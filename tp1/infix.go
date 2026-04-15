package tp1

import "tp1/tdas/pila"

const precedenciaSumaResta = 1
const precedenciaProductoDivision = 2
const precedenciaPotencia = 3

func obtenerPrecedencia(operacion string) int {
	switch operacion {
	case "^":
		return precedenciaPotencia
	case "*", "/":
		return precedenciaProductoDivision
	case "+", "-":
		return precedenciaSumaResta
	default:
		return 0
	}
}

func asociarPorDerecha(operacion string) bool {
	return operacion == "^"
}

func desapilarElemento(tope, actual string) bool {
	if tope == "(" {
		return false
	}
	precedenciaTope := obtenerPrecedencia(tope)
	precedenciaActual := obtenerPrecedencia(actual)

	desapilar := precedenciaTope > precedenciaActual || (precedenciaTope == precedenciaActual && !asociarPorDerecha(actual))

	return desapilar
}

func InfijaAPosfija(tokens []string) []string {
	var salida []string
	pilaConOperadores := pila.CrearPilaDinamica[string]()

	for _, token := range tokens {
		switch token {

		case "+", "-", "*", "/", "^":
			for !pilaConOperadores.EstaVacia() && desapilarElemento(pilaConOperadores.VerTope(), token) {
				salida = append(salida, pilaConOperadores.Desapilar())
			}
			pilaConOperadores.Apilar(token)

		case "(":
			pilaConOperadores.Apilar(token)

		case ")":
			for !pilaConOperadores.EstaVacia() && pilaConOperadores.VerTope() != "(" {
				salida = append(salida, pilaConOperadores.Desapilar())
			}
			pilaConOperadores.Desapilar()
		default:
			salida = append(salida, token)
		}
	}
	for !pilaConOperadores.EstaVacia() {
		salida = append(salida, pilaConOperadores.Desapilar())
	}
	return salida
}
