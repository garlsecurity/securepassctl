package radius

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
			Usage:       "delete a RADIUS",
			ArgsUsage:   "APP_ID",
			Description: "Delete a RADIUS from SecurePass.",
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
		log.Fatal("error: must specify an IP address")
	}
	radius := c.Args()[0]

	if !c.Bool("yes") {
		var reply string
		fmt.Fprintf(os.Stderr,
			"Do you want to delete the RADIUS %q? [y/N] ", radius)
		fmt.Scanln(&reply)
		reply = strings.ToLower(reply)
		if reply != "y" && reply != "yes" {
			os.Exit(-1)
		}
	}

	if _, err := service.Service.RadiusDel(radius); err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("RADIUS deleted: %q", radius)
}
