package main

import (
	"flag"
	"fmt"

	"gitlab.com/joseluiscd/goirs"
)

func main() {
	var file string

	flag.StringVar(&file, "f", "freq.index", "Index file to show")
	flag.Parse()

	findex := goirs.DeserializeFrequencyIndex(file)
	if findex == nil {
		panic("UEUEUEU")
	}

	var ii int
	for _ = range findex.TokenIds {
		ii++
	}
	fmt.Println(ii, "tokens únicos")

	var id int
	for _ = range findex.DocIds {
		id++
	}
	fmt.Println(id, "documentos")

}
