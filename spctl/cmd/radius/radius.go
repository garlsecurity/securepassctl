package radius

import "github.com/urfave/cli"

// Command holds the radius subcommands
var Command = cli.Command{
	Name:        "radius",
	Usage:       "manage RADIUS information",
	ArgsUsage:   "",
	Description: "Manage RADIUS configuration.",
	Subcommands: []cli.Command{},
}
