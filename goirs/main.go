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
	"io"
)

func dieOn(err error){
	if err!=nil {
		panic(err)
	}
}

func main() {
	var (
		tokenize bool
		recordStats bool
		configLoc string

	)

	flag.BoolVar(&tokenize, "tok", false, "Usar sólo el tokenizador")
	flag.StringVar(&configLoc, "config", "./conf.data", "Especifica el archivo de configuración")
	flag.BoolVar(&recordStats, "stats", false, "Especifica si se deben guardar estadísticas")
	flag.Parse()

	config, err := goirs.LoadConfiguration(configLoc)
	dieOn(err)

	dir, err := ioutil.ReadDir(config.Corpus)
	dieOn(err)
	for _, file := range(dir){
		if file.Mode().IsRegular() && strings.HasSuffix(file.Name(), ".html"){
			//Tenemos un fichero candidato
			if tokenize{
				source := filepath.Join(config.Corpus, file.Name())
				dest := filepath.Join(config.Filtered, file.Name()+".ind")
				stats := filepath.Join(config.Stats, "tokenizer.txt")
				tokenizeFile(source, dest, stats)
			}
		}
	}

}

func tokenizeFile(path string, dest string, stats io.Writer){
	ntoks := 0
	file, err := os.Open(path)
	dieOn(err)
	defer file.Close()

	out, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, 0600)
	dieOn(err)
	defer out.Close()

	readfile := bufio.NewReader(file)
	buffer := bytes.NewBuffer(nil)
	toTokenize := bufio.NewWriter(buffer)
	goirs.Filter(readfile, toTokenize)

	t := bufio.NewReader(buffer)
	it := goirs.TokenizerIterator(t)

	for x:= range it{
		ntoks ++
		out.WriteString(x)
		out.WriteString("\n")
	}

}
