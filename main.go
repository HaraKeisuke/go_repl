package main

import (
	"bufio"
	"fmt"
	"go/token"
	"go/types"
	"os"
)

func main() {
	fmt.Printf("Hello Golang!\n\n")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		result, err := types.Eval(token.NewFileSet(), nil, token.NoPos, text)

		if err != nil {
			fmt.Printf("Error: %s", err)
		}

		fmt.Printf(" %v\n\n", result.Value)
	}
}
