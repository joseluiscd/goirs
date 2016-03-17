package goirs

import(
	"sync"
	"encoding/gob"
	"os"
)

//FrequencyIndex es donde se almacenan todos los datos relativos al índice
//de frecuencias
type FrequencyIndex struct {
	//TokenIds guarda el ID de cada token (en forma de string)
	TokenIds map[string]int

	//DocIds guarda el ID de cada documento
	DocIds map[string]int

	TokenNames map[int]string
	DocNames map[int]string
	//NextToken guarda el siguiente número de token que se va a utilizar
	NextToken int

	//NextDoc guarda el siguiente número de documento
	NextDoc int

	//TokensCount es el índice invertido:
	//tokens[token][documento] = veces que aparece token en "documento"
	//En la segunda vuelta esto se divide entre el número máximo de tokens de cada documento
	//En caso de que no esté, siempre se devolverá 0
	//Guardamos esta estructura en caso de que aparezcan nuevos documentos
	//en la colección, para no tener que reconstruirlo todo desde cero.
	TokensCount map[int]map[int]float64

	//Número de veces que aparece el token que más veces aparece en un documento
	MaxTokensDoc map[int]int

	//IDF
	Idfi map[int]float64

	//Pesos
	W map[int]map[int]float64


	mutex sync.Mutex
}

func (ind *FrequencyIndex) AddToken(token string) int {
	if a := ind.TokenIds[token]; a == 0 {
		ind.TokenIds[token] = ind.NextToken
		ind.TokenNames[ind.NextToken] = token
		ind.NextToken++
		return ind.NextToken-1
	} else {
		return a
	}
}

func (ind *FrequencyIndex) AddDocument(doc string) int {
	if a := ind.DocIds[doc]; a == 0 {
		ind.DocIds[doc] = ind.NextDoc
		ind.DocNames[ind.NextDoc] = doc
		ind.NextDoc++
		return ind.NextDoc-1
	} else {
		return a
	}
}

func (ind *FrequencyIndex) AddAndCountToken(doc, token string) {
	ind.mutex.Lock()
	defer ind.mutex.Unlock()

	idToken := ind.AddToken(token)
	idDoc := ind.AddDocument(doc)

	docInd := ind.TokensCount[idToken]

	if docInd == nil {
		docInd = make(map[int]float64)
		docInd[idDoc] = 1
	} else {
		docInd[idDoc]++
	}

	ind.TokensCount[idDoc] = docInd

}

func NewFrequencyIndex() *FrequencyIndex {
	return &FrequencyIndex{
		make(map[string]int),
		make(map[string]int),
		make(map[int]string),
		make(map[int]string),
		1,
		1,
		make(map[int]map[int]float64),
		make(map[int]int),
		make(map[int]float64),
		make(map[int]map[int]float64),
		sync.Mutex{},
	}
}

func (ind* FrequencyIndex) Serialize(file string) {
	stream, err := os.Create(file)
	defer stream.Close()
	if err != nil {
		panic(err)
	}

	encoder := gob.NewEncoder(stream)
	encoder.Encode(ind)
}

func DeserializeFrequencyIndex(file string) *FrequencyIndex {
	stream, err := os.Open(file)
	defer stream.Close()
	if err != nil {
		panic(err)
	}

	toRet := new(FrequencyIndex)

	decoder := gob.NewDecoder(stream)
	decoder.Decode(toRet)

	return toRet
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
