package main

import (
	"os"

	"gitlab.4pd.io/openaios/openaios-iam/cmd/app"
)

func main() {
	command := app.NewIAMCommand()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}

}
