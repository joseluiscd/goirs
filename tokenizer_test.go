package goirs

import(
    "testing"
    "strings"
    "bufio"
)

func TestTokenizer(t *testing.T){
    str := "Esté és el ñtexto.que 10.5  Ño  .debemos déé. . limpiáŕ"
    read := bufio.NewReader(strings.NewReader(str))
    it := TokenizerIterator(read)
    <-it
}
