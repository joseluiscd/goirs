// This file is part of GoIRS.
//
//    GoIRS is free software: you can redistribute it and/or modify
//    it under the terms of the GNU General Public License as published by
//    the Free Software Foundation, either version 3 of the License, or
//    (at your option) any later version.
//
//    GoIRS is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//    GNU General Public License for more details.
//
//    You should have received a copy of the GNU General Public License
//    along with GoIRS.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/joseluiscd/goirs"
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

	fmt.Print("Bienvenido al shell interactivo de GoIRS!!\n(Versión con pesado OKAPI BM25)\n\nGoIRS (OKAPI)-> ")
	for bio.Scan() {
		query := goirs.TokenizerIterator(strings.NewReader(bio.Text())).StopperIterator(stopper).StemmerIterator().ToQuery(index)
		res := goirs.GetOkapiWeight(query, index, config.Okapi.Threshold, config.Okapi.K1, config.Okapi.K3, config.Okapi.B).GetNGreatest()
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
		fmt.Print("GoIRS (OKAPI)-> ")
	}
}
