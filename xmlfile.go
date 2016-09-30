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
