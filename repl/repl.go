package repl

import (
	"bufio"
	"fmt"
	"os"

	"../lexer"
	"../parser"
)

const PROMPT = ">>> "

func Start() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(PROMPT)

		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		lex := lexer.New(line)
		pars := parser.New(lex)

		program := pars.ParseProgram()

		if pars.HasErrors() {
			for _, err := range pars.Errors() {
				fmt.Println("Error: " + err)
			}
			continue
		}

		fmt.Println(program.String())
	}
}
