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
