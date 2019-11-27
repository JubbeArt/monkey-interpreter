package main

import (
	"./repl"
)

func main() {
	//testString := `
	//	let a : true;
	//	return false;
	//	fn (a, b, c) { false };
	//	add(asdf, 2, 3 + 1);
	//
	//`
	//
	//lex := lexer.New(testString)
	//pars := parser.New(lex)
	//syntaxTree := pars.ParseProgram()
	//
	//for _, statement := range syntaxTree.Statements {
	//	switch stmt := statement.(type) {
	//	case fmt.Stringer:
	//		fmt.Println(stmt)
	//		fmt.Println()
	//	default:
	//		fmt.Println("implement string for this statement type")
	//	}
	//}
	//
	//if pars.HasErrors() {
	//	fmt.Println("Errors while parsing: ")
	//	for _, err := range pars.Errors() {
	//		fmt.Println(err)
	//	}
	//}

	repl.Start()

}
