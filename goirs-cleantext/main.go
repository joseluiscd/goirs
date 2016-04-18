package main

// Programa que utiliza funciones de la librería GoIRS para limpiar y tokenizar
// texto de la entrada estándar.

import (
	"fmt"
	"os"

	"joseluiscd/goirs"
)

func main() {
	for x := range goirs.TokenizerIterator(os.Stdin) {
		fmt.Println(goirs.CleanToken(x))
	}
}
