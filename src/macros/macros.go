package macros

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const (
	MacrosFileLimit = 10000
	MacrosLimit     = 10000
)

type MacrosProcessor struct {
	Definitions map[string][]rune
}

func NewMacrosProcessor(filename string) (MacrosProcessor, error) {
	links := make(map[string]string)

	// open file
	f, err := os.Open(filename)
	if err != nil {
		return MacrosProcessor{}, fmt.Errorf("could not open file %s: %v", filename, err)
	}
	macros := make([]byte, MacrosFileLimit)
	n, err := f.Read(macros)
	if err != nil {
		return MacrosProcessor{}, fmt.Errorf("could not read file %s: %v", filename, err)
	}
	if n == 0 {
		return MacrosProcessor{}, fmt.Errorf("file %s is empty", filename)
	}
	if err := json.Unmarshal(macros[:n], &links); err != nil {
		return MacrosProcessor{}, fmt.Errorf("could not unmarshal macros: %v", err)
	}

	// go through links and read files from them
	defs := make(map[string][]rune)

	for k, v := range links {
		f, err := os.Open(v)
		if err != nil {
			return MacrosProcessor{}, fmt.Errorf("could not open file %s: %v", v, err)
		}
		b := make([]byte, MacrosLimit)
		n, err := f.Read(b)
		if err != nil {
			return MacrosProcessor{}, fmt.Errorf("could not read file %s: %v", v, err)
		}
		defs[k] = []rune(string(b[:n]))
	}

	return MacrosProcessor{defs}, nil
}

func (mp MacrosProcessor) Process(src []byte) []byte {
	// TODO: very inefficient, fix it
	for k, v := range mp.Definitions {
		src = []byte(strings.ReplaceAll(string(src), k, string(v)))
	}
	return src
}
