package app

import "github.com/codegangsta/cli"

// Command holds the app subcommands
var Command = cli.Command{
	Name:        "app",
	Usage:       "manage applications",
	ArgsUsage:   "",
	Description: "Manage SecurePass applications.",
	Subcommands: []cli.Command{},
}
