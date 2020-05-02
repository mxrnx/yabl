package main

import (
	"github.com/knarka/yabl/desugarer"
	"github.com/knarka/yabl/parser"
	"github.com/knarka/yabl/tokenizer"
	"io/ioutil"
	"log"
	"os"
)

func usage() {
	log.Fatal("usage: yabl [filename]")
}

func getFileContents(fn string) string {
	if _, err := os.Stat(fn); err != nil {
		log.Println("error: could not find file " + fn)
		usage()
	}

	file, err := os.Open(fn)
	if err != nil {
		log.Println("error: could not open file " + fn)
		usage()
	}
	defer file.Close()


	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("error: could not read file " + fn)
		usage()
	}

	return string(bytes)
}

func main() {
	if len(os.Args) == 1 {
		usage()
		return
	}

	fn := os.Args[1]

	log.Printf("%#v", desugarer.Desugar(parser.Parse(tokenizer.Tokenize(getFileContents(fn)))))
}
