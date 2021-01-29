package main

import (
	"bufio"
	"fmt"
	"os"
	"valley/lexer"
)

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		if text == "" {
			break
		}
		l := lexer.NewLexer(text)
		tokens, err := l.Lex()
		if len(err.Raw) != 0 {
			fmt.Println(err)
			continue
		}
		fmt.Println(tokens)
	}
}
