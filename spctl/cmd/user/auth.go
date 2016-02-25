package user

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/garlsecurity/securepassctl/spctl/service"
)

func init() {
	Command.Subcommands = append(Command.Subcommands,
		cli.Command{
			Name:        "auth",
			Usage:       "test user's authentication",
			ArgsUsage:   "USERNAME PASSWORD",
			Description: "Test user authentication on SecurePass.",
			Action:      ActionAuth,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "no-output, o",
					Usage: "Suppress output",
				},
			},
		})
}

// ActionAuth provides the app list command
func ActionAuth(c *cli.Context) {
	if len(c.Args()) < 2 {
		log.Fatal("error: too few arguments")
	}
	if len(c.Args()) > 2 {
		log.Fatal("error: too many arguments")
	}
	user, password := c.Args()[0], c.Args()[1]

	resp, err := service.Service.UserAuth(user, password)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if !resp.Authenticated {
		if !c.Bool("no-output") {
			fmt.Println("Access denied.")
		}
		os.Exit(-1)
	}
	if !c.Bool("no-output") {
		fmt.Println("Authenticated!")
	}
}
