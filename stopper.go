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
	"bufio"
	"io"
	"os"
)

//Stopper ...
type Stopper map[string]bool

func stopper(stop Stopper, input <-chan string, output chan string) {
	defer close(output)
	for currstr := range input {
		if !stop[currstr] {
			output <- currstr
		}
	}
}

//StopperIterator devuelve aquellos tokens que no están en el stopper
func (tokens StringIterator) StopperIterator(stop Stopper) StringIterator {
	stopped := make(chan string, BUFFERSIZE)

	go stopper(stop, tokens, stopped)

	return stopped
}

//ReadStopper carga un stopper desde una entrada
func ReadStopper(input io.Reader) Stopper {
	stop := make(Stopper)

	scanlines := bufio.NewScanner(input)
	scanlines.Split(bufio.ScanLines)

	for scanlines.Scan() {
		token := CleanToken(scanlines.Text())
		stop[token] = true
	}

	return stop
}

//WriteStopper guarda el stopper (ya normalizado)
func (stopper Stopper) WriteStopper(output io.Writer) {
	for x := range stopper {
		io.WriteString(output, x)
		io.WriteString(output, "\n")
	}
}

//StopperWriterIterator igual que StopperIterator, pero escribe los cambios y se ejecuta opcionalmente
func (tokens StringIterator) StopperWriterIterator(dostop bool, file string, writeStop bool, stop Stopper) StringIterator {
	if dostop {
		if writeStop {
			towrite, err := os.Create(file)
			defer towrite.Close()
			if err != nil {
				panic(err)
			}
			out := make(chan string, BUFFERSIZE)
			in := tokens.StopperIterator(stop)

			dest, err := os.Create(file)

			if err != nil {
				panic(err)
			}
			go tokenWrite(dest, in, out)
			return out

		}
		return tokens.StopperIterator(stop)
	}
	return tokens
}
