package main

import (
	"github.com/dohodges/gofunge/gofunge"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("specify a program file")
	}

	programFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer programFile.Close()

	vm := gofunge.NewVirtualMachine(gofunge.Befunge)
	if err = vm.LoadProgram(programFile); err != nil {
		log.Fatal(err)
	}

	if err = vm.Run(); err != nil {
		log.Fatal(err)
	}
}
