package app

import (
	"log"

	"github.com/codegangsta/cli"
	"github.com/garlsecurity/go-securepass/securepass"
	"github.com/garlsecurity/go-securepass/securepass/spctl/config"
)

func init() {
	Command.Subcommands = append(Command.Subcommands,
		cli.Command{
			Name:        "del",
			Usage:       "delete application",
			ArgsUsage:   "APP_ID",
			Description: "Delete an application from SecurePass.",
			Action:      ActionDel,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "yes, y",
					Usage: "Skip confirmation",
				},
			},
		})
}

// ActionDel provides the del subcommand
func ActionDel(c *cli.Context) {
	if len(c.Args()) != 1 {
		log.Fatal("error: must specify an app id")
	}
	app := c.Args()[0]

	s, err := securepass.NewSecurePass(config.Configuration.AppID,
		config.Configuration.AppSecret, config.Configuration.Endpoint)
	if err != nil {
		log.Fatal(err)
	}
	_, err = s.AppDel(app)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("App deleted: '%s'", app)
}
