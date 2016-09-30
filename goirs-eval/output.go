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
	"encoding/csv"
	"io"
	"strconv"
)

//NewResult crea una línea para el fichero de salida
func NewResult(queryID int, document string) []string {
	return []string{strconv.Itoa(queryID), "0", document, "1"}
}

//CsvEncode guarda en CSV los valores del resultado
func CsvEncode(data [][]string, w io.Writer) {
	writer := csv.NewWriter(w)
	writer.Comma = '\t'

	writer.WriteAll(data)
}
