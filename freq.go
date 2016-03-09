package goirs

import(
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type FrequencyIndex struct {
	//tokenIds guarda el ID de cada token (en forma de string)
	TokenIds map[string]int
	//nextId guarda el siguiente número de token que se va a utilizar
	NextId int
	//tokens es el índice invertido:
	//tokens["documento"][token] = veces que aparece token en "documento"
	Tokens map[string]map[int]int

}

func (ind *FrequencyIndex) AddToken(token string) int {
	if a := ind.TokenIds[token]; a == 0 {
		ind.TokenIds[token] = ind.NextId
		ind.NextId++
		return ind.NextId-1
	} else {
		return a
	}
}

func (ind *FrequencyIndex) AddAndCountToken(doc, token string) {
	idToken := ind.AddToken(token)

	docInd := ind.Tokens[doc]

	if docInd == nil {
		docInd = make(map[int]int)
		docInd[idToken] = 0
	}

	docInd[idToken]++
	ind.Tokens[doc] = docInd

}

func NewFrequencyIndex() *FrequencyIndex {
	a := make(map[string]int)
	b := make(map[string]map[int]int)
	return &FrequencyIndex{a, 1, b}
}

func (ind* FrequencyIndex) Serialize(file string) {
	data, err := json.Marshal(ind.TokenIds)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(file+".tokens", data, 0600)

	tokens := make(map[string]map[string]int)
	for x, y := range ind.Tokens {
		tokens[x] = make(map[string]int)
		for z, w := range y {
			tokens[x][strconv.Itoa(z)] = w
		}
	}

	data, err = json.Marshal(tokens)
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile(file+".freqs", data, 0600)

}

func (tokens StringIterator) IterateFrequencyIndex(document string, index *FrequencyIndex) *FrequencyIndex{
	for x := range tokens {
		index.AddAndCountToken(document, x)
	}
	return index
}

func (tokens StringIterator) AddToFrequencyIndex(doindex bool, document string, index *FrequencyIndex) *FrequencyIndex{
	if doindex {
		return tokens.IterateFrequencyIndex(document, index)
	}
	tokens.Evaluate()
	return nil
}
