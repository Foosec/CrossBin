package main

import (
	"flag"
	"io"
	"log"
	"os"
)

func main() {

	paths := map[string]*string{
		"@binary_lin@": flag.String("l", "", "Linux executable path"),
		"@binary_win@": flag.String("w", "", "Windows executable path"),
		"@binary_mac@": flag.String("m", "", "MacOS executable path"),
	}

	pOut := flag.String("o", "CrossBin.ps1", "Output file")

	flag.Parse()

	packer := Packer{
		PlacePath: paths,
	}
	packer.Init()

	f, err := os.Create(*pOut)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(f, &packer)
}
