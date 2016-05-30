package member

import (
	"log"

	"github.com/codegangsta/cli"
	"github.com/garlsecurity/securepassctl/spctl/service"
)

func init() {
	Command.Subcommands = append(Command.Subcommands,
		cli.Command{
			Name:        "add",
			Usage:       "add group",
			ArgsUsage:   "USERNAME GROUP",
			Description: "Add a user to a group.",
			Action:      ActionAdd,
			Flags:       []cli.Flag{},
		})
}

// ActionAdd provides the add subcommand
func ActionAdd(c *cli.Context) {
	if len(c.Args()) != 2 {
		log.Fatal("error: must specify a username and a group")
	}

	username := c.Args()[0]
	group := c.Args()[1]

	if c.Bool("debug") {
		log.Printf("Adding user %s to group %s\n", username, group)
	}

	_, err := service.Service.GroupMemberAdd(username, group)

	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
