package user

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"
	"github.com/garlsecurity/securepassctl/spctl/service"
)

func init() {
	Command.Subcommands = append(Command.Subcommands,
		cli.Command{
			Name:        "del",
			Usage:       "delete user",
			ArgsUsage:   "USERNAME",
			Description: "Delete a user from SecurePass.",
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
		log.Fatal("error: must specify a username")
	}
	username := c.Args()[0]

	if !c.Bool("yes") {
		var reply string
		fmt.Fprintf(os.Stderr, "Do you want to delete the user %q? [y/N] ", username)
		fmt.Scanln(&reply)
		reply = strings.ToLower(reply)
		if reply != "y" && reply != "yes" {
			os.Exit(-1)
		}
	}

	if _, err := service.Service.UserDel(username); err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("User deleted: %s", username)

}
