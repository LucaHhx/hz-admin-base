package main

import (
	"fmt"
	"os"

	"hab/utils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run ./cmd/hashpw/main.go <password>")
		os.Exit(1)
	}
	fmt.Print(utils.BcryptHash(os.Args[1]))
}
