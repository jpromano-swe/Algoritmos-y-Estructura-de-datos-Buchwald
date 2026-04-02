package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tp0/ejercicios"
)

func main() {
	vector1 := []int{}
	vector2 := []int{}
	var vectorMayor []int

	aperturaArchivo1, _ := os.Open("archivo1.in")
	aperturaArchivo2, _ := os.Open("archivo2.in")
	defer aperturaArchivo1.Close()
	defer aperturaArchivo2.Close()

	scanner1 := bufio.NewScanner(aperturaArchivo1)
	scanner2 := bufio.NewScanner(aperturaArchivo2)

	for scanner1.Scan() {
		ArchivoLimpio1 := scanner1.Text()
		numero, _ := strconv.Atoi(ArchivoLimpio1)
		vector1 = append(vector1, numero)
	}
	for scanner2.Scan() {
		ArchivoLimpio2 := scanner2.Text()
		numero, _ := strconv.Atoi(ArchivoLimpio2)
		vector2 = append(vector2, numero)
	}

	if ejercicios.Comparar(vector1, vector2) >= 0 {
		vectorMayor = vector1
	} else {
		vectorMayor = vector2
	}
	ejercicios.Seleccion(vectorMayor)

	for _, numero := range vectorMayor {
		fmt.Println(numero)
	}
}
