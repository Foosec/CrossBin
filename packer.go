package main

import (
	"bytes"
	"encoding/base64"
	"io"
	"log"
	"os"
	"sort"
)

type Packer struct {
	PlacePath map[string]*string

	Chunks   []io.Reader
	curChunk int
	curPos   int

	script []byte
}

func (p *Packer) Init() error {

	sb, err := os.ReadFile("script.template")
	if err != nil {
		return err
	}
	p.script = sb

	plocs := make(map[int]string)
	//Split by all placeholder
	for k, _ := range p.PlacePath {
		i := bytes.Index(p.script, []byte(k))
		plocs[i] = k
	}

	var indexes []int
	for k, _ := range plocs {
		indexes = append(indexes, k)
	}
	sort.Ints(indexes)

	scriptReader := bytes.NewBuffer(p.script)
	offset := 0
	for _, i := range indexes {

		var buf []byte
		buf = make([]byte, i-offset)
		r, err := scriptReader.Read(buf)
		if err != nil {
			log.Println(err)
			break
		}
		p.Chunks = append(p.Chunks, bytes.NewBuffer(buf))

		fp := *p.PlacePath[plocs[i]]
		if fp != "" {
			f, err := os.Open(fp)
			if err != nil {
				log.Println(err)
			}

			pr, pw := io.Pipe()
			bw := base64.NewEncoder(base64.StdEncoding, pw)
			go func() {
				io.Copy(bw, f)
				bw.Close()
				pw.Close()
			}()

			p.Chunks = append(p.Chunks, pr)
		}

		offset += r
		r, _ = scriptReader.Read(make([]byte, len(plocs[i])))
		offset += r
	}
	buf := make([]byte, len(p.script)-offset)
	scriptReader.Read(buf)
	p.Chunks = append(p.Chunks, bytes.NewBuffer(buf))

	return nil
}

func (p *Packer) Read(dst []byte) (int, error) {

	total := 0
	for total < len(dst) {

		r, err := p.Chunks[p.curChunk].Read(dst[total:])
		if err == io.EOF {
			if p.curChunk == len(p.Chunks)-1 {
				return total + r, io.EOF
			}
			p.curChunk += 1
		} else if err != nil {
			log.Fatal("error:", err)
		}
		total += r
	}

	return total, nil
}
