package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"joseluiscd/goirs"
)

var (
	stopper goirs.Stopper
)

func proccessQuery(query string) {
	//tokens := goirs.TokenizerIterator(strings.NewReader(query)).StopperIterator(stopper)
}

func main() {
	var configLoc string

	fmt.Println("Cargando configuración...")
	flag.StringVar(&configLoc, "config", "./conf.data", "Especifica el archivo de configuración")
	flag.Parse()

	config, err := goirs.LoadConfiguration(configLoc)
	if err != nil {
		panic(err)
	}

	fmt.Println("Cargando índice...")
	index := goirs.DeserializeFrequencyIndex(config.IndexFile)

	fmt.Println("Cargando stopper...")
	stopperfile, err := os.Open(config.StopperFile)
	if err != nil {
		panic(err)
	}
	stopper := goirs.ReadStopper(stopperfile)

	bio := bufio.NewScanner(os.Stdin)

	fmt.Print("Bienvenido al shell interactivo de GoIRS!!\n\nGoIRS -> ")
	for bio.Scan() {
		query := goirs.TokenizerIterator(strings.NewReader(bio.Text())).StopperIterator(stopper).StemmerIterator().ToQuery(index)
		res := goirs.GetQuerySimilarities(query, index).GetNGreatest()
		fmt.Println("Documentos relevantes:")
		i := 0
		for _, val := range res {
			if i > config.MaxDocuments {
				break
			}
			fmt.Println("Documento", index.DocNames[val.DocID], "\tRanking:", val.Weight)
			i++
		}
		if bio.Err() != nil {
			os.Exit(1)
		}
		fmt.Print("GoIRS -> ")
	}
}
