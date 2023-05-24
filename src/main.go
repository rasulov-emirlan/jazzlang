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
	isRepl    = flag.Bool("repl", false, "Starts the REPL")
	isMacros  = flag.Bool("macros", false, "Will parse macros from macros.json in the current directory")
	isVerbose = flag.Bool("verbose", false, "Will print verbose output")
)

func main() {
	flag.Parse()

	mp := macros.MacrosProcessor{}
	err := error(nil)

	if *isMacros {
		mp, err = macros.NewMacrosProcessor("macros.json", *isVerbose)
		if err != nil {
			fmt.Println("Could not create macros processor due to error:", err)
			os.Exit(1)
		}
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

	if len(os.Args) < 2 {
		fmt.Println("Please provide a file to evaluate")
		os.Exit(1)
	}

	file.EvaluateFile(os.Args[len(os.Args)-1], mp)
}
