package user

import (
	"log"

	"github.com/urfave/cli"
	"github.com/garlsecurity/securepassctl/spctl/service"
)

func init() {
	Command.Subcommands = append(Command.Subcommands,
		cli.Command{
			Name:        "enable",
			Usage:       "enable user",
			ArgsUsage:   "USERNAME",
			Description: "Enable a user from SecurePass.",
			Action:      ActionEnable,
		})
}

// ActionDel provides the del subcommand
func ActionEnable(c *cli.Context) {
	if len(c.Args()) != 1 {
		log.Fatal("error: must specify a username")
	}
	username := c.Args()[0]

	if _, err := service.Service.UserEnable(username); err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Printf("User %s enabled", username)

}
