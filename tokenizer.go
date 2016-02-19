package goirs

import (
	"bufio"
	"io"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"regexp"
	"strings"
    "unicode"
	"os"
)


var (
	notallowed = regexp.MustCompile("[^\\p{L}[:digit:]_-]+")
)

func isMn (r rune) bool {
    return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

//CleanToken elimina caracteres extraños de un token y normaliza los acentos
func CleanToken(oldToken string) string {
	oldToken = strings.Replace(oldToken, "ñ", "*", -1)

	//------------------------------------------------------
	// Aquí comienza un bloque de código copiado de StackOverflow...
	b := make([]byte, len(oldToken))

	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	_, _, e := t.Transform(b, []byte(oldToken), true)
	if e != nil {
		return ""
	}
	//Fin del código de StackOverflow
	//-----------------------------------------------------

	return strings.Replace(strings.ToLower(string(b)), "*", "ñ", -1)
}

func cleanToken(in <-chan string, out chan string) {
    defer close(out)

	for currstr := range in {
		if token := CleanToken(currstr); len(token)>0{
			out<-token
		}
	}
}

func tokenizeWords(in <-chan string, out chan string) {
    defer close(out)
	for currstr := range in {
		currstr = notallowed.ReplaceAllString(currstr, " ")
        for _,x := range(strings.Split(currstr, " ")){
            if len(x)>0{
                out <- x
            }
        }
	}
}

func tokenizeSpaces(in *bufio.Scanner, out chan string) {
    defer close(out)
	for in.Scan() {
		currstr := in.Text()
		out <- currstr
	}
}

//TokenizerIterator devuelve un canal que suelta tokens...
func TokenizerIterator(input io.Reader) <-chan string {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)

	uno := make(chan string, 128)
    dos := make(chan string, 128)
    tres := make(chan string, 128)

	go tokenizeSpaces(scanner, uno)
    go tokenizeWords(uno, dos)
    go cleanToken(dos, tres)

	return tres
}

func tokenWrite(file io.Writer, in <-chan string, out chan string) {
	defer close(out)
	for token := range in{
		io.WriteString(file, token)
		io.WriteString(file ,"\n")
		out <- token
	}
}
//TokenizerWriterIterator Igual que TokenizerIterator, pero que escribe los
//tokens en el fichero especificado si write es true
func TokenizerWriterIterator(input io.Reader, file string, write bool) StringIterator {
	if write {
		out := make(chan string, 128)
		in := TokenizerIterator(input)

		dest, err := os.Create(file)
		defer dest.Close()

		if err != nil {
			panic(err)
		}

		go tokenWrite(dest, in, out)

		return out
	}
	return TokenizerIterator(input)
}
