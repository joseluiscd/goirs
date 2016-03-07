package goirs

import(
  "sync"
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
  if a := ind.TokenIds[token]; a != 0 {
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
  docInd[idToken]++
  ind.Tokens[doc] = docInd

}

func (tokens StringIterator) IterateFrequencyIndex(document string, index *FrequencyIndex) *FrequencyIndex{
  for x := tokens {
    index.AddAndCountToken(document, x)
  }
}

func (tokens StringIterator) AddToFrequencyIndex(doindex bool, document string, index *FrequencyIndex) *FrequencyIndex{
  if doindex {
    return tokens.IterateFrequencyIndex()
  } else {
    tokens.Evaluate()
    return nil
  }

}
