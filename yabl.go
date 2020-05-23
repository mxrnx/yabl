package main

import (
	"bufio"
	"fmt"
	"github.com/knarka/yabl/interpreter"
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

func repl() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("o` ")
		code, _ := reader.ReadString('\n')
		fmt.Println(interpreter.Interpret(code).Pretty())
	}
}

func main() {
	if len(os.Args) == 1 {
		usage()
		return
	}

	fn := os.Args[1]

	if fn == "--repl" {
		repl()
	}

	log.Printf("%#v", interpreter.Interpret(getFileContents(fn)))
}
