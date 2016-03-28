package main

import (
	"encoding/xml"
)

//Topic representa una consulta en el XML
type Topic struct {
	XMLName xml.Name `xml:"topic"`
	ID      int      `xml:"id"`
	Desc    string   `xml:"desc"`
}

//Topics representa el conjunto de consultas del XML
type Topics struct {
	XMLName xml.Name `xml:"topics"`
	Topics  []Topic  `xml:",any"`
}
