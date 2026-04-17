package main

import "errors"

type operacion struct {
	simbolo           string
	aridad            int
	asociarPorDerecha bool
	precedencia       int
	operar            func(operadores []int64) (int64, error)
}

func esOperacion(simbolo string) bool {
	_, operadorValido := obtenerOperacion(simbolo)
	return operadorValido
}

func sumar(operadores []int64) (int64, error) {
	return operadores[0] + operadores[1], nil
}
func restar(operadores []int64) (int64, error) {
	return operadores[0] - operadores[1], nil
}

func multiplicar(operadores []int64) (int64, error) {
	return operadores[0] * operadores[1], nil
}

func dividir(operadores []int64) (int64, error) {
	if operadores[1] == 0 {
		return 0, errors.New("No se puede dividir por cero")
	}
	return operadores[0] / operadores[1], nil
}

func potenciar(operadores []int64) (int64, error) {
	resultado := int64(1)

	for i := int64(0); i < operadores[1]; i++ {
		resultado *= operadores[0]
	}
	return resultado, nil
}

func obtenerOperacion(simbolo string) (operacion, bool) {
	switch simbolo {
	case "+":
		return operacion{simbolo: "+", aridad: 2, precedencia: 1, asociarPorDerecha: false, operar: sumar}, true
	case "-":
		return operacion{simbolo: "-", aridad: 2, precedencia: 1, asociarPorDerecha: false, operar: restar}, true
	case "*":
		return operacion{simbolo: "*", aridad: 2, precedencia: 2, asociarPorDerecha: false, operar: multiplicar}, true
	case "/":
		return operacion{simbolo: "/", aridad: 2, precedencia: 2, asociarPorDerecha: false, operar: dividir}, true
	case "^":
		return operacion{simbolo: "^", aridad: 2, precedencia: 3, asociarPorDerecha: true, operar: potenciar}, true
	default:
		return operacion{}, false
	}
}

func debeDesapilarOperador(tope, actual string) bool {
	if tope == "(" {
		return false
	}
	operacionTope, _ := obtenerOperacion(tope)
	operacionActual, _ := obtenerOperacion(actual)

	desapilar := operacionTope.precedencia > operacionActual.precedencia || (operacionTope.precedencia == operacionActual.precedencia && !operacionActual.asociarPorDerecha)

	return desapilar
}
