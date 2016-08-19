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

var (
	k1 = 1.2
	k3 = 1.2
	b  = 0.75
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
func GetOkapiWeight(q Query, ind *FrequencyIndex, thres float64, k1 float64, k3 float64, b float64) QueryResult {
	weights := GetQuerySimilarities(q, ind).GetNGreatest()
	toRet := make(QueryResult)

	for qi := range q {
		for doc := range ind.Weight[qi] {
			ci := weights.GetCi(ind, thres, qi)
			tfi := float64(ind.TokensCount[qi][doc])
			K := float64(k1 * (1 - b + b*(float64(ind.DocLength[doc])/ind.AvgLength)))
			ci *= ((k1 + 1) * tfi) / (K + tfi)
			ci *= (k3 + 1) / (K + 1)

			toRet[doc] += ci
		}
	}
	return toRet
}
