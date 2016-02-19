package goirs

import(
    "io"
    "bufio"
    "os"
)

//Stopper ...
type Stopper map[string]bool

func stopper(stop Stopper, input <-chan string, output chan string) {
    defer close(output)
    for currstr := range input {
        if ! stop[currstr] {
            output <- currstr
        }
    }
}

//StopperIterator devuelve aquellos tokens que no están en el stopper
func (tokens StringIterator) StopperIterator(stop Stopper) StringIterator {
    stopped := make(chan string, BUFFERSIZE)

    go stopper(stop, tokens, stopped)

    return stopped
}

//ReadStopper carga un stopper desde una entrada
func ReadStopper(input io.Reader) Stopper {
    stop := make(Stopper)

    scanlines := bufio.NewScanner(input)
    scanlines.Split(bufio.ScanLines)

    for scanlines.Scan() {
        token := CleanToken(scanlines.Text())
        stop[token] = true
    }

    return stop
}

//WriteStopper guarda el stopper (ya normalizado)
func (stopper Stopper) WriteStopper(output io.Writer){
    for x := range stopper {
        io.WriteString(output, x)
        io.WriteString(output, "\n")
    }
}

//StopperWriterIterator igual que StopperIterator, pero escribe los cambios y se ejecuta opcionalmente
func (tokens StringIterator) StopperWriterIterator(dostop bool, file string, writeStop bool, stop Stopper) StringIterator {
    if dostop {
        if writeStop {
            towrite, err := os.Create(file)
            defer towrite.Close()
            if err != nil {
                panic(err)
            }
            out := make(chan string, BUFFERSIZE)
    		in := tokens.StopperIterator(stop)

    		dest, err := os.Create(file)

    		if err != nil {
    			panic(err)
    		}
    		go tokenWrite(dest, in, out)
    		return out

        }
        return tokens.StopperIterator(stop)
    }
    return tokens
}
