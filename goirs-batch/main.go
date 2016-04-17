package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"gitlab.com/joseluiscd/goirs"
)

var (
	stopper goirs.Stopper
)

func proccessQuery(query string) {
	//tokens := goirs.TokenizerIterator(strings.NewReader(query)).StopperIterator(stopper)
}

func main() {
	var configLoc string
	flag.StringVar(&configLoc, "config", "./conf.data", "Especifica el archivo de configuración")
	flag.Parse()

	//Cargamos la configuración
	config, err := goirs.LoadConfiguration(configLoc)
	if err != nil {
		panic(err)
	}

	//Cargamos el índice
	index := goirs.DeserializeFrequencyIndex(config.IndexFile)

	//Cargamos nuestro stopper
	stopperfile, err := os.Open(config.StopperFile)
	if err != nil {
		panic(err)
	}
	stopper := goirs.ReadStopper(stopperfile)

	//Cargamos nuestro fichero de consultas
	read := goirs.ReadXMLQueries(config)

	for _, d := range read.Topics {
		query := goirs.TokenizerIterator(strings.NewReader(d.Desc)).StopperIterator(stopper).StemmerIterator().ToQuery(index)
		res := goirs.GetQuerySimilarities(query, index).GetNGreatest()

		fmt.Println("\"", d.Desc, "\"")
		i := 1
		for _, val := range res {
			if i >= config.MaxDocuments {
				break
			}

			//Nombre del documento
			fmt.Println("\t", i, ".", index.DocNames[val.DocID])

			//Peso obtenido
			fmt.Println("\t\ta) <", val.Weight, ">")

			//Título del documento
			title, err := goirs.Title(index.DocNames[val.DocID], config)
			if err == nil {
				fmt.Println("\t\tb) <", title, ">")
			} else {
				fmt.Println("Error:", title)
			}

			//Primera frase que contiene una palabra de la consulta
			extract, err := goirs.Extract(query, val.DocID, config, index)
			if err == nil {
				fmt.Println("\t\tc) <", extract, ">")
			} else {
				fmt.Println("Error: ", extract)
			}
			i++
		}
	}
}
