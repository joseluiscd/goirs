package main

import (
	"bufio"
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
	fmt.Println("Cargando índice...")
	index := goirs.DeserializeFrequencyIndex("freq.index")
	fmt.Println("Índice cargado!")
	bio := bufio.NewScanner(os.Stdin)

	fmt.Print("GoIRS -> ")
	for bio.Scan() {

		query := goirs.TokenizerIterator(strings.NewReader(bio.Text())).StopperIterator(stopper).StemmerIterator().ToQuery(index)
		res := goirs.GetQuerySimilarities(query, index).GetNGreatest()
		fmt.Println("Documentos relevantes:")
		i := 0
		for _, val := range res {
			if i > 10 {
				break
			}
			fmt.Println("Documento", index.DocNames[val.DocID], ", Ranking:", val.Weight)
			i++
		}
		if bio.Err() != nil {
			os.Exit(1)
		}
		fmt.Print("GoIRS -> ")
	}
}
