package main

import (
	"log"

	"github.com/Gabriel2233/blog-automator/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("could not execute root command: %s", err)
	}
}
