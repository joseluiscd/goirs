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
	"flag"
	"os"
	"strings"

	"github.com/joseluiscd/goirs"
)

var (
	stopper goirs.Stopper
)

func main() {
	var output [][]string
	var configLoc string
	var okapi bool

	flag.StringVar(&configLoc, "config", "./conf.data", "Especifica el archivo de configuración")
	flag.BoolVar(&okapi, "okapi", false, "Utilizar pesado okapi")
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
		var res goirs.DocumentWeights
		if okapi {
			res = goirs.GetOkapiWeight(query, index, config.Okapi.Threshold, config.Okapi.K1, config.Okapi.K3, config.Okapi.B).GetNGreatest()
		} else {
			res = goirs.GetQuerySimilarities(query, index).GetNGreatest()
		}

		i := 0
		for _, val := range res {
			i++
			r := NewResult(d.ID, index.DocNames[val.DocID])
			output = append(output, r)
			if i == config.MaxDocuments {
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
