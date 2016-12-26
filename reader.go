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
)

func lineReader(scanner *bufio.Scanner, output chan string) {
	for scanner.Scan() {
		line := scanner.Text()
		output <- line
	}
}

//ReadLines reads lines from a file and returns a StringIterator
func ReadLines(inputStream io.Reader) StringIterator {
	scanner := bufio.NewScanner(inputStream)
	scanner.Split(bufio.ScanLines)

	output := make(chan string, BUFFERSIZE)

	go lineReader(scanner, output)
	return output

}
