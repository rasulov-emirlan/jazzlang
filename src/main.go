package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"

	"github.com/rasulov-emirlan/sunjar/src/file"
	"github.com/rasulov-emirlan/sunjar/src/repl"
)

var isRepl = flag.Bool("repl", false, "Starts the REPL")

func main() {
	flag.Parse()

	if *isRepl {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		fmt.Printf(
			"Hello %s! This is the 🎷JAZZ🎷 programming language!\n",
			user.Username,
		)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	}

	if len(os.Args) != 2 {
		fmt.Println("Please provide a file to evaluate")
		os.Exit(1)
	}

	file.EvaluateFile(os.Args[1])
}
