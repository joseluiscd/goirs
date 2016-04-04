package goirs

import "fmt"

//Query representa una consulta en forma de "vector"
type Query map[int]float64

//QueryResult es el resultado de realizar una consulta
type QueryResult map[int]float64

type DocumentWeight struct {
	DocID  int
	Weight float64
}

//GetQuerySimilarities calcula el valor de similitud para todos los documentos
//indexados en ind con respecto a la consulta q
func GetQuerySimilarities(q Query, ind *FrequencyIndex) QueryResult {
	res := make(QueryResult)
	fmt.Println("Entrando")
	for token, wq := range q {
		fmt.Println("UEUE", token)
		for doc, wd := range ind.Weight[token] {
			fmt.Println("Documento", ind.DocNames[doc], "término", ind.TokenNames[token], "peso", wd)
			res[doc] += wq * wd
		}
	}
	return res
}

//GetNGreatest elige los N mejores resultados de una consulta.
//Utiliza un algoritmo semi-quicksort
func (qr QueryResult) GetNGreatest(n int) []DocumentWeight {

	return nil
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
