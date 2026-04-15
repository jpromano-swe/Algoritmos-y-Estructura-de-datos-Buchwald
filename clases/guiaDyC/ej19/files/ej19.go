package files

func Raiz(f func(int) int, a int, b int) int {
	medio := (a + b) / 2

	if f(medio) == 0 {
		return medio
	}

	if (f(medio) < 0 && f(a) < 0) || (f(medio) > 0 && f(a) > 0) {
		return Raiz(f, medio+1, b)
	}
	return Raiz(f, a, medio-1)
}
