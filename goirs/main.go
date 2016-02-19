package main

import (
	"gitlab.com/joseluiscd/goirs"
	"flag"
	"io/ioutil"
	"strings"
	"path/filepath"
	"os"
	"bufio"
	"bytes"
	"fmt"
)

func dieOn(err error){
	if err!=nil {
		panic(err)
	}
}

func main() {
	var (
		configLoc string
		generateConfig bool

		writeTokenized = false
		writeStopped = false

		tokenize = true
		stop = false

		practice = 0
	)

	flag.BoolVar(&generateConfig, "genconfig", false, "Generar un fichero de configuración en el directorio actual")
	flag.StringVar(&configLoc, "config", "./conf.data", "Especifica el archivo de configuración")
	flag.IntVar(&practice, "prac", 1, "Práctica que se quiere ejecutar")
	flag.Parse()

	//Generar configuración si es necesario
	if generateConfig {
		cfg := new(goirs.Configuration)
		err := cfg.Save("./conf.data")
		if err != nil{
			fmt.Println("Fallo al crear la configuración")
			os.Exit(1)
		}
		os.Exit(0)
	}

	//Cargar configuración
	config, err := goirs.LoadConfiguration(configLoc)
	dieOn(err)

	//Prácticas a ejecutar
	switch practice {
	case 2:
		stop = true
		fallthrough
	case 1:
		tokenize = true

	}
	//Decidir qué hacemos en función de la configuración
	if config.Filtered != "" {
		writeTokenized = true
	}

	if config.Stopped != "" {
		writeStopped = true
	}

	//Leer ficheros del corpus y aplicarles las operaciones
	dir, err := ioutil.ReadDir(config.Corpus)
	dieOn(err)

	var tokenized string
	var stopped string

	for _, file := range(dir){
		if file.Mode().IsRegular() && strings.HasSuffix(file.Name(), ".html"){
			//Tenemos un fichero candidato

			/*
			if tokenize{

				dest :=
				tokenizeFile(source, dest)
			}*/
			source := filepath.Join(config.Corpus, file.Name())

			tokenized = filepath.Join(config.Filtered, file.Name()+".tok")
			stopped = filepath.Join(config.Filtered, file.Name()+".tok.stop")

			parsed := goirs.FilterFile(source)

			goirs.TokenizerWriterIterator(parsed, tokenized, writeTokenized)
		}
	}

}

func tokenizeFile(path string, dest string){
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

	for x:= range it{
		ntoks ++
		out.WriteString(x)
		out.WriteString("\n")
	}

}
