package main

import (
	"bufio"
	"gitlab.com/joseluiscd/goirs"
	"os"
)

func main() {
	iii, err := os.Open(os.Args[1])
	if err != nil {
		panic("ueueueue")
	}

	input := bufio.NewReader(iii)
	output := bufio.NewWriter(os.Stdout)

	err = goirs.Filter(input, output)
	if err != nil {
		panic("ueueue")
	}

}
