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

import "math"

//ComputeMaxTokensDoc calcula el campo de la estructura del mismo nombre.
func (ind *FrequencyIndex) ComputeMaxTokensDoc() *FrequencyIndex {
	for _, token := range ind.TokensCount {
		for doc, freq := range token {
			if freq > float64(ind.MaxTokensDoc[doc]) {
				ind.MaxTokensDoc[doc] = int(freq)
			}
		}
	}
	return ind
}

//NormalizeTf divide los TF entre los MaxTokensDoc
func (ind *FrequencyIndex) NormalizeTf() *FrequencyIndex {
	for tokenid, token := range ind.TokensCount {
		for doc, freq := range token {
			ind.TokensCount[tokenid][doc] = freq / float64(ind.MaxTokensDoc[doc])
		}
	}
	return ind
}

//ComputeIdf calcula el IDF de cada token
func (ind *FrequencyIndex) ComputeIdf() *FrequencyIndex {
	ndocs := float64(len(ind.DocIds))
	for token, docs := range ind.TokensCount {
		ind.Idfi[token] = math.Log2(ndocs / float64(len(docs)))
	}
	return ind
}

//ComputeWeights calcula el TF*IDF
func (ind *FrequencyIndex) ComputeWeights() *FrequencyIndex {
	for token, docs := range ind.TokensCount {
		a := ind.Weight[token]
		if a == nil {
			a = make(map[int]float64)
		}

		for doc, tf := range docs {
			//TF * IDF
			a[doc] = tf * ind.Idfi[token]
		}
		ind.Weight[token] = a
	}
	return ind
}

//NormalizeWeights normaliza los pesos
func (ind *FrequencyIndex) NormalizeWeights() *FrequencyIndex {
	for token, docs := range ind.Weight {
		var sum float64
		a := ind.Weight[token]
		for _, w := range docs {
			sum += (w * w)
		}

		n := math.Sqrt(sum)

		for doc, w := range docs {
			if w == 0 || n == 0 {
				a[doc] = 0
			} else {
				a[doc] = w / n
			}
		}

		ind.Weight[token] = a
	}

	return ind
}

//ComputeAverageDocumentLength calcula la longitud media de los documentos
func (ind *FrequencyIndex) ComputeAverageDocumentLength() *FrequencyIndex {
	sum := 0
	ndocs := 0
	for _, l := range ind.DocLength {
		sum += l
		ndocs++
	}

	ind.AvgLength = float64(sum) / float64(ndocs)

	return ind
}

//ComputeAll hace los siguientes cálculos:
// - Tf normalizado (contando la frecuencia máxima de un documento)
// - Idf de cada token
// - Peso normalizado
func (ind *FrequencyIndex) ComputeAll() *FrequencyIndex {
	return ind.ComputeMaxTokensDoc().NormalizeTf().ComputeIdf().ComputeWeights().NormalizeWeights().ComputeAverageDocumentLength()
}
