package radius

import "github.com/codegangsta/cli"

// Command holds the app subcommands
var Command = cli.Command{
	Name:        "radius",
	Usage:       "manage RADIUS information",
	ArgsUsage:   "",
	Description: "Manage RADIUS configuration.",
	Subcommands: []cli.Command{},
}
