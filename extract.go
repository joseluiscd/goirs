package goirs

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

//Title devuelve el título de un documento ya procesado
func Title(doc string, config *Configuration) (string, error) {
	document, err := os.Open(filepath.Join(config.Filtered, doc))
	if err != nil {
		return "<No document>", err
	}

	sc := bufio.NewScanner(document)
	sc.Scan()
	return sc.Text(), nil
}

//Extract extrae de un documento una frase donde aparezca algún término de la consulta
func Extract(query Query, docID int, config *Configuration, index *FrequencyIndex) (string, error) {
	doc := index.DocNames[docID]

	document, err := ioutil.ReadFile(filepath.Join(config.Filtered, doc))
	if err != nil {
		return "<Error 1>", err
	}
	regex := "("
	for x := range query {
		regex += index.TokenNames[x] + "|"
	}

	regex += "$)"
	matcher, err := regexp.Compile(regex)
	if err != nil {
		return "<Error 2>", err
	}

	result := matcher.FindIndex(document)
	a, b := result[0], result[1]

	if a < 0 || b < 0 {
		return "(Not found)", nil
	}

	a -= config.Context
	b += config.Context

	if a < 0 {
		a = 0
	} else {
		for ; a > 0 && document[a] != ' '; a-- {
		}
		if document[a] == ' ' {
			a++
		}
	}

	if b >= len(document) {
		b = len(document) - 1
	} else {
		for ; b <= len(document) && document[b] != ' '; b++ {
		}
	}
	return string(document[a:b]), nil
}
