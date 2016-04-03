package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
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
	data, err := ioutil.ReadFile("test.xml")
	if err != nil {
		panic(err)
	}

	read := Topics{}
	index := goirs.DeserializeFrequencyIndex("freq.index")

	err = xml.Unmarshal(data, &read)
	if err != nil {
		panic(err)
	}
	fmt.Println(index.Weight[8])
	for _, d := range read.Topics {
		query := goirs.TokenizerIterator(strings.NewReader(d.Desc)).StopperIterator(stopper).StemmerIterator().ToQuery(index)
		fmt.Println(query, "UEUEUEUEUEUEUE", goirs.GetQuerySimilarities(query, index))
	}
}
