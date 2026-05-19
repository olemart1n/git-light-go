package cmd

import (
	"fmt"
	"os"
)

func Init() {
	os.Mkdir(".git-light", 0755)
	os.Mkdir(".git-light/objects", 0755)
	os.WriteFile(".git-light/HEAD", []byte(""), 0644)
	fmt.Println("Initialized empty git-light repository")
}
