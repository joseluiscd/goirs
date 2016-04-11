package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"gitlab.com/joseluiscd/goirs"
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
		writeIndex     = false

		stop = false
		stem = false
		freq = false

		practice  = 0
		stopper   goirs.Stopper
		freqindex *goirs.FrequencyIndex
	)

	flag.BoolVar(&generateConfig, "genconfig", false, "Generar un fichero de configuración en el directorio actual")
	flag.StringVar(&configLoc, "config", "./conf.data", "Especifica el archivo de configuración")
	flag.IntVar(&practice, "prac", 4, "Práctica que se quiere ejecutar")
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
	case 4:
		fmt.Print("creación de índice de frecuencias, ")
		freq = true
		freqindex = goirs.NewFrequencyIndex()
		fallthrough
	case 3:
		fmt.Print("stemmer, ")
		stem = true
		fallthrough
	case 2:
		fmt.Print("stopper, ")
		stop = true

		//Cargamos el stopper
		var file *os.File

		if config.StopperFile == "" {
			file, err = os.Open("stopper.txt")
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

	if config.IndexFile != "" {
		writeIndex = true
		fmt.Println("Vamos a escribir el índice de frecuencias (si hay)")
	}

	//Leer ficheros del corpus y aplicarles las operaciones
	dir, err := ioutil.ReadDir(config.Corpus)
	dieOn(err)

	var wg sync.WaitGroup

	worker := func(files <-chan os.FileInfo) {
		var tokenized string
		var stopped string
		var stemmed string

		defer wg.Done()

		for file := range files {
			source := filepath.Join(config.Corpus, file.Name())
			docname := strings.SplitN(file.Name(), ".", 2)[0]

			if writeTokenized {
				tokenized = filepath.Join(config.Filtered, docname+".tok")
			}
			if writeStopped {
				stopped = filepath.Join(config.Stopped, docname+".stop")
			}
			if writeStemmed {
				stemmed = filepath.Join(config.Stemmed, docname+".stem")
			}

			parsed := goirs.FilterFile(source)


			goirs.TokenizerWriterIterator(parsed, tokenized, writeTokenized).
				StopperWriterIterator(stop, stopped, writeStopped, stopper).
				StemmerWriterIterator(stem, stemmed, writeStemmed).
				AddToFrequencyIndex(freq, docname, freqindex)
		}

	}

	files := make(chan os.FileInfo)
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go worker(files)
	}

	for _, file := range dir {
		if file.Mode().IsRegular() && strings.HasSuffix(file.Name(), ".html") {
			//Tenemos un fichero candidato
			files <- file
		}

	}
	close(files)
	wg.Wait()

	if freq {
		freqindex = freqindex.ComputeAll()

		if writeIndex {
			freqindex.Serialize(config.IndexFile)
		}
	}

}
