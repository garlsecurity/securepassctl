package app

import (
	"github.com/codegangsta/cli"
)

func init() {
	Command.Subcommands = append(Command.Subcommands,
		cli.Command{
			Name:        "list",
			Usage:       "list SecurePass's applications",
			ArgsUsage:   " ",
			Description: "List SecurePass's applications.",
			Action:      ActionList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "realm, r",
					Usage: "Set alternate realm",
				},
				cli.BoolFlag{
					Name:  "details, d",
					Usage: "Show more details",
				},
			},
		})
}

func ActionList(c *cli.Context) {

}
