package main

import (
	"fmt"
	"monkey/reader"
	"monkey/repl"
	"os"
	"os/user"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		checkError(err)
		defer file.Close()

		if strings.Contains(file.Name(), ".monkey") {
			reader.Start(file, os.Stdout)
		}
	} else {
		user, err := user.Current()
		checkError(err)

		fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
		fmt.Printf("Feel free to type in commands\n")
		fmt.Printf("Type 'exit' to quit.\n")
		repl.Start(os.Stdin, os.Stdout)
	}
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
