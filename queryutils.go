package goirs

type Query struct {
	Weight map[int]float64
}

func GetQuerySimilarity(ind FrequencyIndex) {

}

func (query Query) normalizeQuery(ind FrequencyIndex) Query {

	return query
}

func (tokens StringIterator) ToQuery(ind FrequencyIndex) Query {
	var q Query

	for token := range tokens {
		id := ind.TokenIds[token]
		if id != 0 {
			q.Weight[id]++
		}
	}
	return q.normalizeQuery(ind)
}
