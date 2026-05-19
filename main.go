package main

import (
	"fmt"
	"git-light/cmd"

	"os"
)

func main() {

	arguments := os.Args // Args (array) hold the command-line arguments, starting with the program name.
	if len(arguments) < 2 {
		fmt.Print("No arguments provided")
		return
	}

	command := os.Args[1]

	switch command {
	case "init":
		cmd.Init()
	case "commit": // git-light commit -m "message.."
		if len(arguments) < 4 || arguments[2] != "-m" {
			fmt.Println("usage: git-light commit -m <message>")
			return
		}
		cmd.Commit(".", arguments[3])

	case "log":
		cmd.Log()

	default:
		fmt.Println("git-light: unknown command:", command)

	}
}
