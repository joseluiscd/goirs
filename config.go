package goirs

import (
	"encoding/json"
	"io/ioutil"
)

//Configuration es la estructura en la que guardamos la configuración
type Configuration struct {
	Corpus   string
	Filtered string
	Stopped  string
	Stats string "."
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
