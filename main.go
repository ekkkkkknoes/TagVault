package main

import (
	"fmt"
	"os"

	"github.com/ekkkkkknoes/TagVault/cmd"
)

func main() {
	app := cmd.CreateApp()
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
