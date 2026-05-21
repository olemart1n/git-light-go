package main

import (
	"fmt"
	"git-light/cmd"
	"git-light/repo"

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
		err := cmd.Commit(".", arguments[3])
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
		// UNDER ER OUTPUT FRA EKTE GIT - KANSKJE FÅ TIL NOE TILSVARENDE SENERE
		// 		mart@air git-light-go % git commit -m "log funksjonalitet"
		// [main 4a9888d] log funksjonalitet
		//  4 files changed, 117 insertions(+), 3 deletions(-)
		//  create mode 100644 README.md
		//  create mode 100644 repo/parse-commit.go

	case "log":
		cmd.Log()

	case "checkout":
		if len(arguments) < 3 {
			fmt.Println("Usage: git-light checkout <hash>")
			return
		}
		repo.Checkout(arguments[2])
	default:
		fmt.Println("git-light: unknown command:", command)

	}
}
