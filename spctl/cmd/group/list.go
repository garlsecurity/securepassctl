package group

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
			Usage:       "list SecurePass's groups",
			ArgsUsage:   " ",
			Description: "list SecurePass's groups.",
			Action:      ActionGroupList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "noheaders, H",
					Usage: "Don't print headers",
				},
				cli.StringFlag{
					Name:  "realm, r",
					Usage: "Set alternate realm",
				},
			},
		})
}

// ActionGroupList provides the group list command
func ActionGroupList(c *cli.Context) {
	if len(c.Args()) != 0 {
		log.Fatal("too many arguments")
	}

	resp, err := service.Service.GroupList(c.String("realm"))
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if !c.Bool("noheaders") {
		log.Printf("%s\n", "GROUP")
	}

	for _, user := range resp.Group {
		fmt.Println(user)
	}
}
