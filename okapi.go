package goirs

const (
	k1 float64 = 1.2
	k3 float64 = 1.2
	b  float64 = 0.75
)

//GetCi devuelve el peso Robertson-Sparck Jones de un término dada una consulta
func (k DocumentWeights) GetCi(ind *FrequencyIndex, thres float64, token int) float64 {
	var V float64
	var Vi float64

	for ds := range k {
		if k[ds].Weight > thres {
			V++
			if ind.Weight[token][k[ds].DocID] > 0 {
				Vi++
			}
		}
	}
	dfi := float64(len(ind.Weight[token]))
	N := float64(ind.NextDoc - 1)
	return ((Vi + 0.5) / (V - Vi + 0.5)) / ((dfi - Vi + 0.5) / (N - dfi - V + Vi + 0.5))
}

//GetOkapiWeight devuelve el peso Okapi BM25 para la consulta dada
func GetOkapiWeight(q Query, ind *FrequencyIndex) QueryResult {
	weights := GetQuerySimilarities(q, ind).GetNGreatest()
	toRet := make(QueryResult)

	for qi := range q {
		for doc := range ind.Weight[qi] {
			ci := weights.GetCi(ind, 0.5, qi)
			tfi := float64(ind.TokensCount[qi][doc])
			K := float64(k1 * (1 - b + b*(float64(ind.DocLength[doc])/ind.AvgLength)))
			ci *= ((k1 + 1) * tfi) / (K + tfi)
			ci *= (k3 + 1) / (K + 1)

			toRet[doc] += ci
		}
	}
	return toRet
}
