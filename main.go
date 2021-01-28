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
		fmt.Println(l.Lex())
	}
}
