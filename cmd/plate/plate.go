package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/neelr/plate/pkg/create"
	"github.com/neelr/plate/pkg/load"
)

func main() {
	flagsCmd := flag.NewFlagSet("create", flag.ExitOnError)
	flagsName := flagsCmd.String("name", "", "Name for new template")

	if len(os.Args) < 2 {
		// Do a help
		fmt.Println("expected create")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create":
		flagsCmd.Parse(os.Args[2:])
		create.Command(*flagsName)
	case "load":
		flagsCmd.Parse(os.Args[2:])
		load.Command(*flagsName)
	}
}
