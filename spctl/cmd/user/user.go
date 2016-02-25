package user

import "github.com/codegangsta/cli"

// Command holds the user subcommands
var Command = cli.Command{
	Name:        "user",
	Usage:       "manage users",
	ArgsUsage:   "",
	Description: "Manage SecurePass users.",
	Subcommands: []cli.Command{},
}
