package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"

	"github.com/rasulov-emirlan/sunjar/src/file"
	"github.com/rasulov-emirlan/sunjar/src/macros"
	"github.com/rasulov-emirlan/sunjar/src/repl"
)

var (
	isRepl     = flag.Bool("repl", false, "Starts the REPL")
	macrosFlag = flag.String("macros", "macros.json", "Path to the macros file")
)

func main() {
	flag.Parse()

	mp, err := macros.NewMacrosProcessor(*macrosFlag)
	if err != nil {
		fmt.Println("Could not create macros processor due to error:", err)
		os.Exit(1)
	}

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

	file.EvaluateFile(os.Args[1], mp)
}
