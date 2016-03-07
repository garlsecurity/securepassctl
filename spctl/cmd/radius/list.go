package radius

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
			Usage:       "list SecurePass's RADIUS devices",
			ArgsUsage:   " ",
			Description: "List SecurePass's RADIUS devices.",
			Action:      ActionList,
			Flags: []cli.Flag{
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

// ActionList provides the radius list command
func ActionList(c *cli.Context) {
	if len(c.Args()) != 0 {
		log.Fatal("too many arguments")
	}

	resp, err := service.Service.RadiusList(c.String("realm"))
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if c.Bool("details") {
		fmt.Printf("%-45s %-30s\n", "RADIUS", "FQDN")
	}

	for _, radius := range resp.IPAddrs {
		if !c.Bool("details") {
			fmt.Printf("%s\n", radius)
		} else {
			r, e := service.Service.RadiusInfo(radius)
			if e != nil {
				log.Fatalf("error: couldn't retrieve details for %q: %v",
					radius, e)
			}
			fmt.Printf("%-45s %-30s\n", radius, r.Name)
		}
	}
}
