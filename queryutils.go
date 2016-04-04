package goirs

import (
	"fmt"
	"sort"
)

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
	fmt.Println("Entrando")
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
