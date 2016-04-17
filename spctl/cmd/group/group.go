package group

import "github.com/codegangsta/cli"

// Command holds the user subcommands
var Command = cli.Command{
	Name:        "group",
	Usage:       "manage groups",
	ArgsUsage:   "",
	Description: "Manage SecurePass groups.",
	Subcommands: []cli.Command{},
}
