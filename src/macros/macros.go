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

// MacrosItem is a struct for json unmarshalling.
// Important: SrcPlain and SrcFile are mutually exclusive.
type MacrosItem struct {
	SrcPlain *string `json:"src_plain"`
	SrcFile  *string `json:"src_file"`
}

type MacrosProcessor struct {
	Definitions map[string][]rune
}

func NewMacrosProcessor(filename string) (MacrosProcessor, error) {
	links := make(map[string]MacrosItem)

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
		if v.SrcPlain == nil && v.SrcFile == nil {
			return MacrosProcessor{}, fmt.Errorf("neither src_plain nor src_file is defined for %s", k)
		}
		if v.SrcPlain != nil && v.SrcFile != nil {
			return MacrosProcessor{}, fmt.Errorf("both src_plain and src_file are defined for %s, they are mutually exclusive", k)
		}

		if v.SrcPlain != nil {
			defs[k] = []rune(*v.SrcPlain)
			continue
		}

		// open file
		f, err := os.Open(string(*v.SrcFile))
		if err != nil {
			return MacrosProcessor{}, fmt.Errorf("could not open file %s: %v", string(*v.SrcFile), err)
		}
		macros := make([]byte, MacrosLimit)
		n, err := f.Read(macros)
		if err != nil {
			return MacrosProcessor{}, fmt.Errorf("could not read file %s: %v", string(*v.SrcFile), err)
		}
		if n == 0 {
			return MacrosProcessor{}, fmt.Errorf("file %s is empty", string(*v.SrcFile))
		}
		defs[k] = []rune(string(macros[:n]))

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
