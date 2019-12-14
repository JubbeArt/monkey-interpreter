package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"../lexer"
	"../tokens"
)

const PROMPT = ">>> "

func Repl() {
	scanner := bufio.NewScanner(os.Stdin)
	//env := object.NewEnvironment()

	for {
		fmt.Print(PROMPT)

		if !scanner.Scan() {
			break
		}

		//line := scanner.Text()
		//lex := lexer.New(line)
		//pars := parser.New(lex)
		//
		//program := pars.ParseProgram()
		//
		//if pars.HasErrors() {
		//	for _, err := range pars.Errors() {
		//		fmt.Println("Error: " + err)
		//	}
		//	continue
		//}
		//
		//result := evaluator.Eval(program, env)
		//
		//if result == nil {
		//	//fmt.Println("eval returned nil")
		//} else {
		//	fmt.Println(result.String())
		//}
	}
}

func Rlpl(input io.Reader) {
	scanner := bufio.NewScanner(input)

	for {
		fmt.Print(PROMPT)

		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		lex := lexer.New(line)

		for {
			token := lex.NextToken()

			if token.Type == tokens.EOF {
				break
			}

			fmt.Println(token)
		}
	}
}
