package main

import (
	"os"

	"github.com/allenhaozi/crabgo/cmd/app"
)

func main() {
	command := app.NewCrabCommand()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}

}
