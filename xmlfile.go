package goirs

import (
	"encoding/xml"
	"io/ioutil"
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

//ReadXMLQueries lee las consultas de un fichero XML
func ReadXMLQueries(config *Configuration) Topics {
	var topics Topics

	data, err := ioutil.ReadFile(config.QueryFile)
	if err != nil {
		panic(err)
	}

	xml.Unmarshal(data, &topics)

	return topics
}
