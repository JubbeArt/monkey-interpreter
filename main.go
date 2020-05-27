package main

import (
	"fmt"

	"./lexer"
	"./parser"
)

func main() {
	//const (
	//	RLPL = iota // lex
	//	RPPL        // parse
	//	REPL        // evalutate
	//)

	//replType := REPL

	//var input io.Reader
	//
	//if len(os.Args) >= 2 {
	//	input, _ = os.Open(os.Args[1])
	//} else {
	//	input = os.Stdin
	//}
	//
	//file, _ := ioutil.ReadFile(os.Args[1])
	//
	//repl.Rlpl(os.Stdin)
	lex := lexer.New(`
	`)
	//
	//for token := lex.NextToken(); token.Type != tokens.EOF; token = lex.NextToken() {
	//	fmt.Println(token)
	//}
	//
	pars := parser.New(lex)
	program := pars.ParseProgram()
	//_ = program
	//
	if pars.HasErrors() {
		pars.PrintErrors()
	} else {
		fmt.Println(program.String(0))
	}

	//if replType == RLPL {
	//	repl.Rlpl()
	//} else if replType == RPPL {
	//	fmt.Println("TODO")
	//} else {
	//	repl.Repl()
	//}
	//f, _ := filepath.Glob("/home/jesper/Pictures/.downloads/*/*")
	//fmt.Println(strings.Join(f, "\n"))
}
