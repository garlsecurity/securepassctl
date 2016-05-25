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
			Name:        "glist",
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
				cli.BoolFlag{
					Name:  "details, d",
					Usage: "Show more details",
				},
			},
		})
}

// ActionGroupList provides the group list command
func ActionGroupList(c *cli.Context) {
	if len(c.Args()) != 0 {
		log.Fatal("too many arguments")
	}

	resp, err := service.Service.UserList(c.String("realm"))
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if c.Bool("details") && !c.Bool("noheaders") {
		log.Printf("%-35s %-35s %s\n", "USERNAME", "FULL NAME", "STATUS")
	}

	for _, user := range resp.Username {
		if !c.Bool("details") {
			fmt.Println(user)
		} else {
			r, e := service.Service.UserInfo(user)
			if e != nil {
				log.Fatalf("couldn't retrieve details for '%s': %s",
					user, err)
			}
			status := "Active"
			if !r.Enabled {
				status = "Disabled"
			}
			fmt.Printf("%-35s %-35s %s\n", user, fmt.Sprintf("%s %s",
				r.Name, r.Surname), status)
		}
	}
}
