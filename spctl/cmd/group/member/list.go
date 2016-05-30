package member

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
	"github.com/garlsecurity/securepassctl/spctl/service"
)

func init() {
	Command.Subcommands = append(Command.Subcommands,
		cli.Command{
			Name:        "list",
			Usage:       "list SecurePass's users in group",
			ArgsUsage:   "GROUP",
			Description: "list SecurePass's users in group.",
			Action:      ActionGroupMemberList,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "noheaders, H",
					Usage: "Don't print headers",
				},
			},
		})
}

// ActionGroupList provides the group list command
func ActionGroupMemberList(c *cli.Context) {
	if len(c.Args()) != 1 {
		log.Fatal("error: specify group")
	}

	group := c.Args()[0]

	resp, err := service.Service.GroupMemberList(group)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if !c.Bool("noheaders") {
		log.Printf("%s\n", "USERNAME")
	}

	for _, user := range resp.Members {
		fmt.Println(user)
	}
}
