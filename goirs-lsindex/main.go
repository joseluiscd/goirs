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
	"fmt"

	"github.com/joseluiscd/goirs"
	"github.com/kljensen/snowball/spanish"
)

func main() {
	var term string
	var indexLoc string

	flag.StringVar(&indexLoc, "index", "./index", "Index location")
	flag.StringVar(&term, "t", "jaén", "Término a buscar")
	flag.Parse()

	findex := goirs.DeserializeFrequencyIndex(indexLoc)
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
