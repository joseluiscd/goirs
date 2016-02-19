package goirs

//StringIterator es un iterador genérico de cadenas
type StringIterator <-chan string

const(
    //BUFFERSIZE es el tamaño del buffer de los canales
    BUFFERSIZE = 64
)
//Next devuelve el siguiente valor o la cadena vacía en caso de llegar al final
func (i StringIterator) Next() string {
    st, ok := <-i
    if ok {
        return st
    }
    return ""
}

func (i StringIterator) filter(f func(string) bool, out chan string){
    defer close(out)
    for x := range i {
        if f(x){
            out<-x
        }
    }
}

//Filter filtra el contenido del iterador dependiendo de la función f
func (i StringIterator) Filter(f func(string) bool) StringIterator {
    k := make(chan string, BUFFERSIZE)
    go i.filter(f, k)
    return k
}

//Evaluate cierra la cadena y evalua todo lo del iterador
//Básicamente, lo saca todo para forzar las operaciones intermedias
func (i StringIterator) Evaluate() {
    for _ = range i {}
}
