package files

func ParteEnteraRaiz(n int) int {
	return _ParteEnteraRaiz(n, 0, n, 0)
}

func _ParteEnteraRaiz(n, ini, fin, resultado int) int {
	if ini > fin {
		return resultado
	}

	medio := (ini + fin) / 2

	if medio*medio == n {
		return medio
	}

	if medio*medio < n {
		return _ParteEnteraRaiz(n, medio+1, fin, medio)
	}

	if medio*medio > n {
		return _ParteEnteraRaiz(n, ini, medio-1, resultado)
	}
	return resultado
}
