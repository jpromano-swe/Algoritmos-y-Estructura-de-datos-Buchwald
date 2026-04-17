package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		linea := scanner.Text()
		tokens := tokenizar(linea)
		resultado := InfijaAPosfija(tokens)
		fmt.Println(strings.Join(resultado, " "))
	}
}
