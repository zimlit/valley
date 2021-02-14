package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		if text == "" {
			break
		}
	}
}
