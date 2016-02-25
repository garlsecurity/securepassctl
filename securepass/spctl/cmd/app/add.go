package app

import (
	"log"

	"github.com/codegangsta/cli"
	"github.com/garlsecurity/securepassctl/securepass"
	"github.com/garlsecurity/securepassctl/securepass/spctl/service"
)

func init() {
	Command.Subcommands = append(Command.Subcommands,
		cli.Command{
			Name:        "add",
			Usage:       "add application to SecurePass",
			ArgsUsage:   "LABEL",
			Description: "Add an application to SecurePass.",
			Action:      ActionAdd,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "ipv4, 4",
					Usage: "restrict to IPv4 network (default: 0.0.0.0/0)",
				},
				cli.StringFlag{
					Name:  "ipv6, 6",
					Usage: "restrict to IPv6 network (default: ::/0)",
				},
				cli.StringFlag{
					Name:  "group, g",
					Usage: "Group name (restriction)",
				},
				cli.StringFlag{
					Name:  "realm, r",
					Usage: "Set alternate realm",
				},
				cli.BoolFlag{
					Name:  "write, w",
					Usage: "Write capabilities (default: read-only)",
				},
				cli.BoolFlag{
					Name:  "privacy, p",
					Usage: "Enable privacy mode (default: false)",
				},
			},
		})
}

// ActionAdd provides the add subcommand
func ActionAdd(c *cli.Context) {
	if len(c.Args()) != 1 {
		log.Fatal("error: must specify a label")
	}
	label := c.Args()[0]

	resp, err := service.Service.AppAdd(&securepass.ApplicationDescriptor{
		Label:            label,
		Group:            c.String("group"),
		Realm:            c.String("realm"),
		Write:            c.Bool("write"),
		Privacy:          c.Bool("privacy"),
		AllowNetworkIPv4: c.String("ipv4"),
		AllowNetworkIPv6: c.String("ipv6"),
	})
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Printf("Application ID: %s\nApplication Secret: %s\n\n"+
		"Please save the above results, secret will be no longer displayed for security reasons.",
		resp.AppID, resp.AppSecret)
}
