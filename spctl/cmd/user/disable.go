package user

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/garlsecurity/securepassctl/spctl/service"
)

func init() {
	Command.Subcommands = append(Command.Subcommands,
		cli.Command{
			Name:        "disable",
			Usage:       "disable user",
			ArgsUsage:   "USERNAME",
			Description: "Disable a user from SecurePass.",
			Action:      ActionDisable,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "yes, y",
					Usage: "Skip confirmation",
				},
			},
		})
}

// ActionDel provides the del subcommand
func ActionDisable(c *cli.Context) {
	if len(c.Args()) != 1 {
		log.Fatal("error: must specify a username")
	}

	username := c.Args()[0]

	if !c.Bool("yes") {
		var reply string
		fmt.Fprintf(os.Stderr, "Do you want to disable user %q? [y/N] ", username)
		fmt.Scanln(&reply)
		reply = strings.ToLower(reply)

		if reply != "y" && reply != "yes" {
			os.Exit(-1)
		}
	}

	if _, err := service.Service.UserDisable(username); err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Printf("User %s deleted", username)

}
