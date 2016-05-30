package member

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
			Name:        "del",
			Usage:       "delete group member",
			ArgsUsage:   "USERNAME GROUP",
			Description: "Delete a user from group.",
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
	if len(c.Args()) != 2 {
		log.Fatal("error: must specify a username and a group")
	}

	username := c.Args()[0]
	group := c.Args()[1]

	if !c.Bool("yes") {
		var reply string

		fmt.Fprintf(os.Stderr, "Do you want to delete the user %q in group %q? [y/N] ", username, group)
		fmt.Scanln(&reply)

		reply = strings.ToLower(reply)

		if reply != "y" && reply != "yes" {
			os.Exit(-1)
		}
	}

	if _, err := service.Service.GroupMemberDel(username, group); err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Printf("User %s deleted from group %s", username, group)

}
