package group

import (
	"log"

	"github.com/codegangsta/cli"
	"github.com/garlsecurity/securepassctl"
	"github.com/garlsecurity/securepassctl/spctl/service"
)

func init() {
	Command.Subcommands = append(Command.Subcommands,
		cli.Command{
			Name:        "add",
			Usage:       "add group",
			ArgsUsage:   "GROUP",
			Description: "Add a group to SecurePass.",
			Action:      ActionAdd,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "description, d",
					Usage: "Description",
				},
			},
		})
}

// ActionAdd provides the add subcommand
func ActionAdd(c *cli.Context) {
	if len(c.Args()) != 1 {
		log.Fatal("error: must specify a group")
	}
	group := c.Args()[0]

	if c.Bool("debug") {
		log.Println("Adding group", group)
	}

	_, err := service.Service.GroupAdd(&securepassctl.GroupDescriptor{
		Group:       group,
		Description: c.String("description"),
	})

	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
