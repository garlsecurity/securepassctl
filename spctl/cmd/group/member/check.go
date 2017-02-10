package member

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
	"github.com/garlsecurity/securepassctl/spctl/service"
)

func init() {
	Command.Subcommands = append(Command.Subcommands,
		cli.Command{
			Name:        "check",
			Usage:       "test group membership",
			ArgsUsage:   "USERNAME GROUP",
			Description: "Check whether a SecurePass user belongs to a group.",
			Action:      ActionGroupMember,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "no-output, o",
					Usage: "Suppress output",
				},
			},
		})
}

// ActionGroupMember handles the group-member command
func ActionGroupMember(c *cli.Context) {
	if len(c.Args()) < 2 {
		log.Fatal("error: too few arguments")
	}
	if len(c.Args()) > 2 {
		log.Fatal("error: too many arguments")
	}
	user, group := c.Args()[0], c.Args()[1]

	resp, err := service.Service.GroupMember(user, group)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if !c.Bool("no-output") {
		msg := fmt.Sprintf("User %s belongs to group %s!", user, group)
		if !resp.Member {
			msg = fmt.Sprintf("User %s not in group %s!", user, group)
		}
		fmt.Println(msg)
	}

	if !resp.Member {
		os.Exit(-1)
	}
}
