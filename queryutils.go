package goirs

type Query map[int]float64

type QueryResult struct {
	Document   int
	Similarity float64
}

type QueryResults []QueryResult

func GetQuerySimilarity(q Query, ind FrequencyIndex) {

}

func (tokens StringIterator) ToQuery(ind FrequencyIndex) Query {
	var q Query

	for token := range tokens {
		id := ind.TokenIds[token]
		if id != 0 {
			q[id]++
		}
	}
	return q
}
