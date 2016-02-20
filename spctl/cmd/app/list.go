package app

import (
	"log"

	"github.com/codegangsta/cli"
	"github.com/garlsecurity/go-securepass/securepass"
	"github.com/garlsecurity/go-securepass/spctl/config"
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

// ActionList provides the app list command
func ActionList(c *cli.Context) {
	if len(c.Args()) != 0 {
		log.Fatal("too many arguments")
	}

	s, err := securepass.NewSecurePass(config.Configuration.AppID,
		config.Configuration.AppSecret, config.Configuration.Endpoint)
	if err != nil {
		log.Fatal(err)
	}

	s.AppList()
}
