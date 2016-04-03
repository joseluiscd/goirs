package goirs

import "fmt"

//Query representa una consulta en forma de "vector"
type Query map[int]float64

//QueryResult es el resultado de realizar una consulta
type QueryResult map[int]float64

//GetQuerySimilarities calcula el valor de similitud para todos los documentos
//indexados en ind con respecto a la consulta q
func GetQuerySimilarities(q Query, ind *FrequencyIndex) QueryResult {
	res := make(QueryResult)
	fmt.Println(ind.Weight)
	for id, doc := range ind.Weight {
		var score float64
		for t, p := range q {
			score += doc[t] * p
		}

		res[id] = score
	}

	return res
}

//ToQuery transforma un iterador de cadenas en el formato requerido para una consulta
func (tokens StringIterator) ToQuery(ind *FrequencyIndex) Query {
	q := make(Query)

	for token := range tokens {
		id := ind.TokenIds[token]
		if id != 0 {
			q[id]++
		}
	}
	return q
}
