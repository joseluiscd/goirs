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
	"encoding/gob"
	"encoding/xml"
	"os"
	"strconv"
	"sync"
)

//FrequencyIndex es donde se almacenan todos los datos relativos al índice
//de frecuencias
type FrequencyIndex struct {
	//TokenIds guarda el ID de cada token (en forma de string)
	TokenIds map[string]int

	//DocIds guarda el ID de cada documento
	DocIds map[string]int

	TokenNames map[int]string
	DocNames   map[int]string
	//NextToken guarda el siguiente número de token que se va a utilizar
	NextToken int

	//NextDoc guarda el siguiente número de documento
	NextDoc int

	//TokensCount es el índice invertido:
	//tokens[token][documento] = veces que aparece token en "documento"
	//En la segunda vuelta esto se divide entre el número máximo de tokens de cada documento
	//En caso de que no esté, siempre se devolverá 0
	//Guardamos esta estructura en caso de que aparezcan nuevos documentos
	//en la colección, para no tener que reconstruirlo todo desde cero.
	TokensCount map[int]map[int]float64

	//Número de veces que aparece el token que más veces aparece en un documento
	MaxTokensDoc map[int]int

	//IDF
	Idfi map[int]float64

	//Pesos
	Weight map[int]map[int]float64

	DocLength map[int]int
	AvgLength float64
	mutex     sync.Mutex
}

//AddToken añade un token al índice de frecuencias
func (ind *FrequencyIndex) AddToken(token string) int {
	a := ind.TokenIds[token]
	if a == 0 {
		ind.TokenIds[token] = ind.NextToken
		ind.TokenNames[ind.NextToken] = token
		ind.NextToken++
		return ind.NextToken - 1
	}
	return a
}

func (ind *FrequencyIndex) AddTokenWithId(token string, id int) {
	ind.TokenIds[token] = id
	ind.TokenNames[id] = token
}

//AddDocument añade (parcialmente) un documento al índice de frecuencias. El cómputo
//de tokens y demás debe realizarse por separado
func (ind *FrequencyIndex) AddDocument(doc string) int {
	a := ind.DocIds[doc]
	if a == 0 {
		ind.DocIds[doc] = ind.NextDoc
		ind.DocNames[ind.NextDoc] = doc
		ind.NextDoc++
		return ind.NextDoc - 1
	}

	return a
}

//AddDocumentWithId añade un documento al índice de frecuencias con el ID dado
//El cómputo de tokens y demás debe realizarse por separado
func (ind *FrequencyIndex) AddDocumentWithId(doc string, id int) {
	ind.DocIds[doc] = id
	ind.DocNames[id] = doc
}

//AddAndCountToken añade el token (si no está) y lo cuenta al documento especificado
func (ind *FrequencyIndex) AddAndCountToken(doc, token string) {
	ind.mutex.Lock()
	defer ind.mutex.Unlock()

	idToken := ind.AddToken(token)
	idDoc := ind.AddDocument(doc)

	ind.DocLength[idDoc]++
	docInd := ind.TokensCount[idToken]

	if docInd == nil {
		docInd = make(map[int]float64)
		docInd[idDoc] = 1
	} else {
		docInd[idDoc]++
	}

	ind.TokensCount[idToken] = docInd
}

//CountToken cuenta un token en el documento indicado
func (ind *FrequencyIndex) CountToken(idDoc, idToken int) {
	ind.mutex.Lock()
	defer ind.mutex.Unlock()

	ind.DocLength[idDoc]++
	docInd := ind.TokensCount[idToken]

	if docInd == nil {
		docInd = make(map[int]float64)
		docInd[idDoc] = 1
	} else {
		docInd[idDoc]++
	}

	ind.TokensCount[idToken] = docInd
}

//NewFrequencyIndex es el constructor del índice de frecuencias
func NewFrequencyIndex() *FrequencyIndex {
	return &FrequencyIndex{
		make(map[string]int),
		make(map[string]int),
		make(map[int]string),
		make(map[int]string),
		1,
		1,
		make(map[int]map[int]float64),
		make(map[int]int),
		make(map[int]float64),
		make(map[int]map[int]float64),
		make(map[int]int),
		0,
		sync.Mutex{},
	}
}

//Serialize serializa el índice de frecuencias a un archivo
func (ind *FrequencyIndex) Serialize(file string) {
	stream, err := os.Create(file)
	defer stream.Close()
	if err != nil {
		panic(err)
	}

	encoder := gob.NewEncoder(stream)
	encoder.Encode(ind)
}

//MarshalXML Exports FrequencyIndex to XML
func (ind *FrequencyIndex) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)

	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "TF"}})
	for key, value := range ind.TokensCount {
		t := xml.StartElement{Name: xml.Name{Local: "term"}, Attr: []xml.Attr{xml.Attr{Name: xml.Name{Local: "Id"}, Value: ind.TokenNames[key]}}}
		tokens := []xml.Token{t}

		for key2, value2 := range value {
			attr := []xml.Attr{xml.Attr{Name: xml.Name{Local: "Id"}, Value: ind.DocNames[key2]}}
			tokens = append(tokens,
				xml.StartElement{Name: xml.Name{Local: "document"}, Attr: attr},
				xml.CharData(strconv.FormatFloat(value2, 'g', -1, 64)),
				xml.EndElement{Name: xml.Name{Local: "document"}})

		}
		tokens = append(tokens, xml.EndElement{Name: xml.Name{Local: "term"}})
		for _, t := range tokens {
			err := e.EncodeToken(t)
			if err != nil {
				return err
			}
		}
	}
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "TF"}})
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "IDF"}})

	for key, value := range ind.Idfi {
		t := xml.StartElement{Name: xml.Name{Local: "term"}, Attr: []xml.Attr{xml.Attr{Name: xml.Name{Local: "Id"}, Value: ind.TokenNames[key]}}}
		tokens := []xml.Token{t}

		tokens = append(tokens,
			xml.CharData(strconv.FormatFloat(value, 'g', -1, 64)),
			xml.EndElement{Name: t.Name})

		for _, t := range tokens {
			err := e.EncodeToken(t)
			if err != nil {
				return err
			}
		}

	}

	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "IDF"}})

	e.EncodeToken(xml.EndElement{Name: start.Name})

	// flush to ensure tokens are written
	err := e.Flush()
	if err != nil {
		return err
	}

	return nil
}

//SerializeXML exports the FrequencyIndex to a XML file
func (ind *FrequencyIndex) SerializeXML(file string) {
	stream, err := os.Create(file)
	defer stream.Close()
	if err != nil {
		panic(err)
	}

	encoder := xml.NewEncoder(stream)
	err = encoder.Encode(ind)
	if err != nil {
		panic(err)
	}
}

//DeserializeFrequencyIndex carga un índice de frecuencias desde un fichero
func DeserializeFrequencyIndex(file string) *FrequencyIndex {
	stream, err := os.Open(file)
	defer stream.Close()
	if err != nil {
		panic(err)
	}

	toRet := new(FrequencyIndex)

	decoder := gob.NewDecoder(stream)
	decoder.Decode(toRet)

	return toRet
}

//IterateFrequencyIndex itera sobre un canal de tokens, añadiéndolos al índice de frecuencias en el documento especificado.
//En este caso, el canal de tokens representa un documento
func (tokens StringIterator) IterateFrequencyIndex(document string, index *FrequencyIndex) *FrequencyIndex {
	for x := range tokens {
		index.AddAndCountToken(document, x)
	}
	return index
}

//AddToFrequencyIndex es el iterador de alto nivel, preparado para el main
func (tokens StringIterator) AddToFrequencyIndex(doindex bool, document string, index *FrequencyIndex) *FrequencyIndex {
	if doindex {
		return tokens.IterateFrequencyIndex(document, index)
	}
	tokens.Evaluate()
	return nil
}
