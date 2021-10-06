package main

import (
	"bytes"
	b64 "encoding/base64"
	"flag"
	"io"
	"log"
	"os"
	"strings"
)

const (
	winPathI = iota
	linPathI = iota
	macPathI = iota
)

func main() {

	paths := make([]string, 3)

	flag.StringVar(&paths[winPathI], "w", "", "Windows executable path")
	flag.StringVar(&paths[linPathI], "l", "", "Linux executable path")
	flag.StringVar(&paths[macPathI], "m", "", "MacOS executable path")

	pOut := flag.String("o", "CrossBin.ps1", "Output file")

	flag.Parse()

	r, err := packBinaries(paths)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(*pOut)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(f, r)
}

func packBinaries(paths []string) (io.Reader, error) {

	templatePlaceholders := []string{"@binary_win@", "@binary_lin@", "@binary_mac@"}

	script, err := os.ReadFile("script.template")
	if err != nil {
		return nil, err
	}

	for i, p := range paths {

		if p == "" {
			script = []byte(strings.Replace(string(script), templatePlaceholders[i], "", 1))
			continue
		}

		bin, err := os.ReadFile(p)
		if err != nil {
			return nil, err
		}
		sbin := b64.StdEncoding.EncodeToString(bin)
		script = []byte(strings.Replace(string(script), templatePlaceholders[i], sbin, 1))
	}

	return bytes.NewReader(script), nil

}
