package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tp1/tp1"
	"unicode"
)

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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		linea := scanner.Text()
		tokens := tokenizar(linea)
		resultado := tp1.InfijaAPosfija(tokens)
		fmt.Println(strings.Join(resultado, " "))
	}
}
