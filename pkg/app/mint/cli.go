package app

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

const (
	AppName    = "mint"
	AppVersion = "0.0.1"
	AppUsage   = "track your spend"
)

func newCLI() *cli.App {

	cliApp := cli.NewApp()

	cliApp.Name = AppName
	cliApp.Version = AppVersion
	cliApp.Usage = AppUsage

	cliApp.CommandNotFound = func(ctx *cli.Context, command string) {
		fmt.Printf("unknown command - %v \n\n", command)
		cli.ShowAppHelp(ctx)
	}

	cliApp.Action = func(c *cli.Context) error {
		fmt.Println("Hello mint user")
		return nil
	}
	return cliApp
}
