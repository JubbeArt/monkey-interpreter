package main

import (
	"fmt"
	"path/filepath"
	"strings"
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
	//repl.Rlpl(input)
	//file, _ := ioutil.ReadFile(os.Args[1])
	//
	//lex := lexer.New(string(file))
	//
	////for token := lex.NextToken(); token.Type != tokens.EOF; token = lex.NextToken() {
	////	fmt.Println(token)
	////}
	//
	//pars := parser.New(lex)
	//program := pars.ParseProgram()
	//
	//if pars.HasErrors() {
	//	pars.PrintErrors()
	//} else {
	//	env := object.NewEnvironment()
	//	evaluator.Eval(program, env)
	//}

	//if replType == RLPL {
	//	repl.Rlpl()
	//} else if replType == RPPL {
	//	fmt.Println("TODO")
	//} else {
	//	repl.Repl()
	//}
	f, _ := filepath.Glob("/home/jesper/Pictures/.downloads/*/*")
	fmt.Println(strings.Join(f, "\n"))
}
