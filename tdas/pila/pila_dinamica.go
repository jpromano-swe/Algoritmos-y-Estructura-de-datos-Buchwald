package pila

/* Definición del struct pila proporcionado por la cátedra. */

const FactorDeCrecimiento = 2

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, 0)
	pila.cantidad = 0
	return pila
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	if pila.cantidad == 0 {
		return true
	}
	return false
}

func (pila *pilaDinamica[T]) VerTope() T {
	if pila.cantidad == 0 {
		panic("La pila esta vacia")
	}

	return pila.datos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) Apilar(elemento T) {
	pila.Redimensionar()

	pila.datos[pila.cantidad] = elemento

	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {

	if pila.cantidad == 0 {
		panic("La pila esta vacia")
	}
	pila.cantidad--

	return pila.datos[pila.cantidad]
}

func (pila *pilaDinamica[T]) Redimensionar() {
	if pila.cantidad == len(pila.datos) {
		nuevaCapacidad := len(pila.datos) * FactorDeCrecimiento
		if nuevaCapacidad == 0 {
			nuevaCapacidad = FactorDeCrecimiento
		}

		nuevosDatos := make([]T, nuevaCapacidad)

		for i := 0; i < pila.cantidad; i++ {
			nuevosDatos[i] = pila.datos[i]
		}
		pila.datos = nuevosDatos
	}
}
