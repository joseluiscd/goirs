package main

import (
	"flag"
	"os"
	"strings"

	"joseluiscd/goirs"
)

var (
	stopper goirs.Stopper
)

func main() {
	var output [][]string
	var configLoc string

	flag.StringVar(&configLoc, "config", "./conf.data", "Especifica el archivo de configuración")
	flag.Parse()

	config, err := goirs.LoadConfiguration(configLoc)
	if err != nil {
		panic(err)
	}

	stopperfile, err := os.Open(config.StopperFile)
	if err != nil {
		panic(err)
	}

	stopper := goirs.ReadStopper(stopperfile)
	read := goirs.ReadXMLQueries(config)
	index := goirs.DeserializeFrequencyIndex(config.IndexFile)

	for _, d := range read.Topics {
		query := goirs.TokenizerIterator(strings.NewReader(d.Desc)).StopperIterator(stopper).StemmerIterator().ToQuery(index)
		res := goirs.GetQuerySimilarities(query, index).GetNGreatest()

		i := 0
		for _, val := range res {
			i++
			r := NewResult(d.ID, index.DocNames[val.DocID])
			output = append(output, r)
			if i == 5 {
				break
			}
		}
	}
	write, err := os.Create(config.EvalFile)
	if err != nil {
		panic(err)
	}

	CsvEncode(output, write)
	write.Close()
}
