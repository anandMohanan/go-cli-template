package main

import (
	"os"
	"github.com/charmbracelet/log"
)

func handlePanic() {
	if err := recover(); err != nil {
		log.Error("Program panicked", err)
		os.Exit(1)
	}
}

func main() {
	defer handlePanic()
    

}
