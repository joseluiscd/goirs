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

import "sort"

//Query representa una consulta en forma de "vector"
type Query map[int]float64

//QueryResult es el resultado de realizar una consulta
type QueryResult map[int]float64

//DocumentWeight representa el valor de similitud de un documento con una consulta
type DocumentWeight struct {
	DocID  int
	Weight float64
}

//DocumentWeights representa una lista de similitudes
type DocumentWeights []DocumentWeight

//GetQuerySimilarities calcula el valor de similitud para todos los documentos
//indexados en ind con respecto a la consulta q
func GetQuerySimilarities(q Query, ind *FrequencyIndex) QueryResult {
	res := make(QueryResult)
	for token, wq := range q {
		for doc, wd := range ind.Weight[token] {
			//fmt.Println("Documento", ind.DocNames[doc], "término", ind.TokenNames[token], "peso", wd)
			res[doc] += wq * wd
		}
	}
	return res
}

func (d DocumentWeights) Len() int {
	return len(d)
}

func (d DocumentWeights) Less(i, j int) bool {
	return d[i].Weight < d[j].Weight
}

func (d DocumentWeights) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

//GetNGreatest elige los N mejores resultados de una consulta.
//Utiliza el sort de todos los valores (ineficiente)
func (qr QueryResult) GetNGreatest() DocumentWeights {
	res := make(DocumentWeights, len(qr))
	i := 0
	for doc, weight := range qr {
		res[i] = DocumentWeight{doc, weight}
		i++
	}

	sort.Sort(sort.Reverse(res))
	return res
}

//ToQuery transforma un iterador de cadenas en el formato requerido para una consulta
func (tokens StringIterator) ToQuery(ind *FrequencyIndex) Query {
	q := make(Query)

	for token := range tokens {
		id := ind.TokenIds[token]
		if id != 0 {
			q[id] = 1
		}
	}
	return q
}
