package main

import (
	"log"

	"github.com/Exar04/orBox/cmd"
)

func main() {
    err := cmd.Execute()
    if err != nil {
        log.Fatal(err)
    }
}
