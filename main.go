package main

import (
	"fmt"
	"github.com/ahmadrosid/yuk/compiler"
	"github.com/ahmadrosid/yuk/lexer"
	"github.com/ahmadrosid/yuk/parser"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatalf("please provide file path!")
	}

	result, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	lex := lexer.New(string(result))
	par := parser.New(lex)
	gen := compiler.New(par.ParseProgram())
	fmt.Println(gen.Generate())
}
