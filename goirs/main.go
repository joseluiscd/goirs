package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"gitlab.com/joseluiscd/goirs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func dieOn(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	var (
		configLoc      string
		generateConfig bool

		writeTokenized = false
		writeStopped   = false
		writeStemmed   = false

		stop = false
		stem = false

		practice = 0
		stopper  goirs.Stopper
	)

	flag.BoolVar(&generateConfig, "genconfig", false, "Generar un fichero de configuración en el directorio actual")
	flag.StringVar(&configLoc, "config", "./conf.data", "Especifica el archivo de configuración")
	flag.IntVar(&practice, "prac", 3, "Práctica que se quiere ejecutar")
	flag.Parse()

	//Generar configuración si es necesario
	if generateConfig {
		cfg := new(goirs.Configuration)
		err := cfg.Save("./conf.data")
		if err != nil {
			fmt.Println("Fallo al crear la configuración")
			os.Exit(1)
		}
		os.Exit(0)
	}

	//Cargar configuración
	config, err := goirs.LoadConfiguration(configLoc)
	dieOn(err)

	//Prácticas a ejecutar
	fmt.Printf("Ejecutando la práctica %d...\nSe va a utilizar: ", practice)
	switch practice {
	case 3:
		fmt.Print("stemmer, ")
		stem = true
		fallthrough
	case 2:
		fmt.Print("stopper, ")
		stop = true

		//Cargamos el stopper
		var file *os.File
		var err error

		if config.StopperFile == "" {
			file, err = os.Open(filepath.Join(config.Index, "stopper.txt"))
			dieOn(err)
		} else {
			file, err = os.Open(config.StopperFile)
			dieOn(err)
		}

		stopper = goirs.ReadStopper(file)
		file.Close()

		fallthrough
	case 1:
		fmt.Println("filtrado y tokenizado")

	}

	//Decidir qué hacemos en función de la configuración
	if config.Filtered != "" {
		writeTokenized = true
		fmt.Println("Vamos a escribir el tokenizado (si hay)")
	}

	if config.Stopped != "" {
		writeStopped = true
		fmt.Println("Vamos a escribir el stopper (si hay)")
	}

	if config.Stemmed != "" {
		writeStemmed = true
		fmt.Println("Vamos a escribir el stemmer (si hay)")
	}

	//Leer ficheros del corpus y aplicarles las operaciones
	dir, err := ioutil.ReadDir(config.Corpus)
	dieOn(err)

	var tokenized string
	var stopped string
	var stemmed string

	for _, file := range dir {
		if file.Mode().IsRegular() && strings.HasSuffix(file.Name(), ".html") {
			//Tenemos un fichero candidato

			source := filepath.Join(config.Corpus, file.Name())

			if writeTokenized {
				tokenized = filepath.Join(config.Filtered, file.Name()+".tok")
			}
			if writeStopped {
				stopped = filepath.Join(config.Filtered, file.Name()+".tok.stop")
			}
			if writeStemmed {
				stemmed = filepath.Join(config.Filtered, file.Name()+".tok.stop.stem")
			}

			parsed := goirs.FilterFile(source)

			goirs.TokenizerWriterIterator(parsed, tokenized, writeTokenized).
				StopperWriterIterator(stop, stopped, writeStopped, stopper).
				StemmerWriterIterator(stem, stemmed, writeStemmed).
				Evaluate()
		}
	}

}

func tokenizeFile(path string, dest string) {
	ntoks := 0
	file, err := os.Open(path)
	dieOn(err)
	defer file.Close()

	out, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, 0600)
	dieOn(err)
	defer out.Close()

	readfile := bufio.NewReader(file)
	buffer := bytes.NewBuffer(nil)

	goirs.Filter(readfile, buffer)

	it := goirs.TokenizerIterator(buffer)

	for x := range it {
		ntoks++
		out.WriteString(x)
		out.WriteString("\n")
	}

}
