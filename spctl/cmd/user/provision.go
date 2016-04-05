package user

import (
	"log"

	"github.com/codegangsta/cli"
	"github.com/garlsecurity/securepassctl/spctl/service"
)

var tokenTypes = map[string]bool{
	"iphone":     true,
	"android":    true,
	"blackberry": true,
	"software":   true,
	"securepass": true,
	"googleauth": true,
}

func init() {
	Command.Subcommands = append(Command.Subcommands,
		cli.Command{
			Name:        "provision",
			Usage:       "token provisioning",
			ArgsUsage:   "USERNAME",
			Description: "Provision user.",
			Action:      ActionProvision,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "token, t",
					Usage: "Token type (securepass, googleauth)",
					Value: "securepass",
				},
			},
		})
}

// ActionProvision provides the provision subcommand
func ActionProvision(c *cli.Context) {
	if len(c.Args()) != 1 {
		log.Fatal("error: must specify a username")
	}

	username := c.Args()[0]
	token := c.String("token")
	if !tokenTypes[token] {
		log.Fatalf("error: invalid token: '%s'", token)
	}

	_, err := service.Service.UserProvision(username, token)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
