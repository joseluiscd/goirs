package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"

	"gitlab.com/joseluiscd/goirs"
)

var (
	stopper goirs.Stopper
)

func proccessQuery(query string) {
	//tokens := goirs.TokenizerIterator(strings.NewReader(query)).StopperIterator(stopper)
}

func main() {
	var a []Topic

	a = append(a, Topic{ID: 1, Desc: "Primooooooo"})
	a = append(a, Topic{ID: 2, Desc: "eueueueueueueu"})

	f, err := xml.MarshalIndent(Topics{Topics: a}, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(f))

	data, err := ioutil.ReadFile("test.xml")
	if err != nil {
		panic(err)
	}

	read := Topics{}

	err = xml.Unmarshal(data, &read)
	if err != nil {
		panic(err)
	}

	for _, d := range read.Topics {
		fmt.Println(d)
	}
}
