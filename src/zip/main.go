package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/mzrimsek/zip-lang/src/zip/reader"
	"github.com/mzrimsek/zip-lang/src/zip/repl"
)

func main() {
	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		checkError(err)
		defer file.Close()

		reader.Start(file, os.Stdout)
	} else {
		user, err := user.Current()
		checkError(err)

		fmt.Printf("Hello %s! This is the Zip programming language!\n", user.Username)
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
