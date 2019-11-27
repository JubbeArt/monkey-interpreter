package main

import (
	"fmt"

	"./lexer"
	"./parser"
)

func main() {
	testString := `
		true;
		if (true) { false };
		if (1) { 123 } else { false };
	`

	lex := lexer.New(testString)
	pars := parser.New(lex)
	syntaxTree := pars.ParseProgram()

	for _, statement := range syntaxTree.Statements {
		switch stmt := statement.(type) {
		case fmt.Stringer:
			fmt.Println(stmt)
		default:
			fmt.Println("implement string for this statement type")
		}
	}

	if pars.HasErrors() {
		fmt.Println("Errors while parsing: ")
		for _, err := range pars.Errors() {
			fmt.Println(err)
		}
	}

}
