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
	for token, docs := range ind.TokensCount {
		ind.Idfi[token] = math.Log2(float64(ind.NextDoc-1) / float64(len(docs)))
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
