package goirs

import(
	"math"
)
//ComputeMaxTokensDoc calcula el campo de la estructura del mismo nombre.
func (ind *FrequencyIndex) ComputeMaxTokensDoc() *FrequencyIndex{
	//TODO: Hacer concurrente si va muy lento
	for _, docs := range ind.TokensCount {
		for doc, freq := range docs {
			if ind.MaxTokensDoc[doc] < freq {
				ind.MaxTokensDoc[doc] = freq
			}
		}
	}

	return ind
}

//NormalizeTf divide los TF entre los MaxTokensDoc
func (ind *FrequencyIndex) NormalizeTf() *FrequencyIndex {
	
	return ind
}

//ComputeIdfi calcula el IDF de cada token
func (ind *FrequencyIndex) ComputeIdfi() *FrequencyIndex {
	for token, docs := range ind.TokensCount {
		ind.Idfi[token] = math.Log2(float64(ind.NextDoc-1)/len(docs))
	}
	return ind
}

func (ind *FrequencyIndex) ComputeWeights() *FrequencyIndex {
	//OJO OJO OJO

	return ind
}
