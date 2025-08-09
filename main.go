package main

import (
	"os"

	"github.com/darshan744/container/cmd"
)

func main() {
	args := os.Args[1:]

	switch args[0] {
	case "run":
		cmd.Run(args[1:])

	case "child":
		cmd.Child()
	}

}
