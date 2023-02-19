package main

import (
	"log"

	"github.com/rcebrian/dotfiles-cli/cmd"
)

func init() {
	log.SetFlags(0)
}

func main() {
	_ = cmd.Execute()
}
