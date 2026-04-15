package main

import (
  "clases/guiaDyC/files"
  "fmt"
)

func main() {
  f := func(x int) int {
    return x - 5
  }

  intervaloA := 0
  intervaloB := 10

  resultado := files.Raiz(f, intervaloA, intervaloB)

  fmt.Printf("Probando f(x) = x - 5 en el intervalo [%d, %d]\n", intervaloA, intervaloB)
  fmt.Printf("Raíz encontrada: %d\n", resultado)
  fmt.Printf("Verificación: f(%d) = %d\n", resultado, f(resultado))
}
