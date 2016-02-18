package goirs

import(
    "io"
    "bufio"
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

//StopperIterator toma como parámetro un canal que devuelve tokens y devuelve
// aquellos que no están en el stopper
func StopperIterator(tokens <-chan string, stop Stopper) chan string {
    stopped := make(chan string, 128)

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
func WriteStopper(stopper Stopper, output io.Writer){
    for x := range stopper {
        io.WriteString(output, x)
        io.WriteString(output, "\n")
    }
}
