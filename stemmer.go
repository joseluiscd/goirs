package goirs

import (
    "os"
    "github.com/kljensen/snowball/spanish"
)

func stemmer(input <-chan string, output chan string){
    defer close(output)
    for currstr := range input {
        output <- spanish.Stem(currstr, false)
    }
}

//StemmerIterator recorre el iterador y devuelve
func (tokens StringIterator) StemmerIterator() StringIterator {
    out := make(chan string, BUFFERSIZE)
    go stemmer(tokens, out)
    return out
}

//StemmerWriterIterator escribe los tokens pasados por el stemmer
func (tokens StringIterator) StemmerWriterIterator(dostem bool, file string, writeStem bool) StringIterator {
    if dostem {
        if writeStem {
            towrite, err := os.Create(file)
            defer towrite.Close()
            if err != nil {
                panic(err)
            }
            out := make(chan string, BUFFERSIZE)
    		in := tokens.StemmerIterator()

    		dest, err := os.Create(file)

    		if err != nil {
    			panic(err)
    		}
    		go tokenWrite(dest, in, out)
    		return out

        }
        return tokens.StemmerIterator()
    }
    return tokens
}
