package goirs

import (
	"encoding/json"
	"io/ioutil"
)

//Configuration es la estructura en la que guardamos la configuración
type Configuration struct {
	Corpus       string
	Filtered     string
	Tokenized    string
	Stopped      string
	StopperFile  string
	Stemmed      string
	IndexFile    string
	QueryFile    string
	EvalFile     string
	Context      int
	MaxDocuments int

	//Peso necesario para considerar un documento relevante al calcular el peso okapi
	Okapi struct {
		Threshold float64
		K1        float64
		K3        float64
		B         float64
	}
}

//LoadConfiguration carga un archivo de configuración
func LoadConfiguration(file string) (*Configuration, error) {
	var toRet *Configuration

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	toRet = new(Configuration)
	err = json.Unmarshal(data, toRet)
	if err != nil {
		return nil, err
	}

	return toRet, nil
}

//Save guarda la configuración en un archivo
func (cfg *Configuration) Save(file string) error {

	content, err := json.MarshalIndent(cfg, "", "\t")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(file, content, 0600)
}
