package main

import (
	"encoding/csv"
	"io"
	"strconv"
)

//Result es un resultado de un documento en una consulta
type Result []string

//NewResult crea una línea para el fichero de salida
func NewResult(queryID int, document string) Result {
	return Result{strconv.Itoa(queryID), "0", document, "1"}
}

//CsvEncode guarda en CSV los valores del resultado
func CsvEncode(data []Result, w io.Writer) {
	writer := csv.NewWriter(w)
	writer.Comma = '\t'
}
