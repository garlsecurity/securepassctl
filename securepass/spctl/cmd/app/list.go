package app

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
	"github.com/garlsecurity/go-securepass/securepass"
	"github.com/garlsecurity/go-securepass/securepass/spctl/config"
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

	resp, err := config.Configuration.AppList(&securepass.ApplicationDescriptor{
		Realm: c.String("realm"),
	})
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if c.Bool("details") {
		fmt.Printf("%-45s %-30s\n", "APP_ID", "LABEL")
	}

	for _, app := range resp.AppID {
		if !c.Bool("details") {
			fmt.Printf("%s\n", app)
		} else {
			r, e := config.Configuration.AppInfo(app)
			if e != nil {
				log.Fatalf("couldn't retrieve details for '%s': %s",
					app, err)
			}
			fmt.Printf("%-45s %-30s\n", app, r.Label)
		}
	}
}
