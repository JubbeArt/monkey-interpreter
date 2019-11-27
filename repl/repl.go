package repl

import (
	"bufio"
	"fmt"
	"os"

	"../lexer"
	"../tokens"
)

const PROMPT = ">>> "

func Start() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(PROMPT)

	for scanner.Scan() {
		line := scanner.Text()
		lex := lexer.New(line)

		for token := lex.NextToken(); token.Type != tokens.EOF; token = lex.NextToken() {
			fmt.Println(token)
		}

		fmt.Print(PROMPT)
	}
}
