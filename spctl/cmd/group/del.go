package group

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
			Usage:       "delete group",
			ArgsUsage:   "GROUP",
			Description: "Delete a group from SecurePass.",
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
		log.Fatal("error: must specify a group")
	}

	group := c.Args()[0]

	if !c.Bool("yes") {
		var reply string

		fmt.Fprintf(os.Stderr, "Do you want to delete the group %q? [y/N] ", group)
		fmt.Scanln(&reply)

		reply = strings.ToLower(reply)

		if reply != "y" && reply != "yes" {
			os.Exit(-1)
		}
	}

	if _, err := service.Service.GroupDel(group); err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Printf("Group deleted: %s", group)

}
