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

package goirs

import (
	"os"

	"github.com/kljensen/snowball/spanish"
)

func stemmer(input <-chan string, output chan string) {
	defer close(output)
	for currstr := range input {
		output <- spanish.Stem(currstr, false)
	}
}

//StemmerIterator recorre el iterador y devuelve
func (tokens StringIterator) StemmerIterator() StringIterator {
	out := make(chan string, BUFFERSIZE)
	go stemmer(tokens, out)
	return out
}

//StemmerWriterIterator escribe los tokens pasados por el stemmer
func (tokens StringIterator) StemmerWriterIterator(dostem bool, file string, writeStem bool) StringIterator {
	if dostem {
		if writeStem {
			towrite, err := os.Create(file)
			defer towrite.Close()
			if err != nil {
				panic(err)
			}
			out := make(chan string, BUFFERSIZE)
			in := tokens.StemmerIterator()

			dest, err := os.Create(file)

			if err != nil {
				panic(err)
			}
			go tokenWrite(dest, in, out)
			return out

		}
		return tokens.StemmerIterator()
	}
	return tokens
}
