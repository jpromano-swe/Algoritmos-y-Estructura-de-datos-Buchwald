package main

import "unicode"

func esOperadorOParentesis(caracter string) bool {
	switch caracter {
	case "+", "-", "*", "/", "^", "(", ")":
		return true
	default:
	}
	return false
}

func tokenizar(linea string) []string {
	tokens := []string{}
	i := 0

	for i < len(linea) {
		caracter := linea[i]

		if unicode.IsSpace(rune(caracter)) {
			i++
		} else if esOperadorOParentesis(string(caracter)) {
			tokens = append(tokens, string(caracter))
			i++
		} else {
			inicio := i
			for i < len(linea) && linea[i] >= '0' && linea[i] <= '9' {
				i++
			}
			tokens = append(tokens, linea[inicio:i])
		}
	}
	return tokens
}
