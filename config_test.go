package goirs

import (
	"testing"
	"fmt"
)

//TestLoadConfig sirve para probar LoadConfiguration
func TestLoadConfig(t *testing.T) {
	cfg, err := LoadConfiguration("config.sample")
	if err != nil {
		t.Fail()
	}

	if cfg.Corpus == "asdf" &&
		cfg.Filtered == "ueue" &&
		cfg.Stopped == "st" {

		fmt.Print("Guay!!")

	} else {
		t.Fail()
	}
}
