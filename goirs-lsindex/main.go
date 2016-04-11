package main

import (
	"flag"
	"fmt"

	"github.com/kljensen/snowball/spanish"
	"gitlab.com/joseluiscd/goirs"
)

func main() {
	var term string
	var configLoc string

	flag.StringVar(&configLoc, "config", "./conf.data", "Especifica el archivo de configuración")
	flag.StringVar(&term, "t", "jaén", "Término a buscar")
	flag.Parse()

	config, err := goirs.LoadConfiguration(configLoc)
	if err != nil {
		panic(err)
	}

	findex := goirs.DeserializeFrequencyIndex(config.IndexFile)
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

	term = spanish.Stem(term, false)
	idterm := findex.TokenIds[term]
	if idterm != 0 {
		fmt.Println("IDF:", findex.Idfi[idterm])
	}



}
