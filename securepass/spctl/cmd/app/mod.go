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
			Name:        "mod",
			Usage:       "modify application",
			ArgsUsage:   "LABEL",
			Description: "Modify an application in SecurePass",
			Action:      ActionMod,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "label, l",
					Usage: "Label",
				},
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
				cli.BoolFlag{
					Name:  "write, w",
					Usage: "Write capabilities (default: false)",
				},
				cli.BoolTFlag{
					Name:  "read, r",
					Usage: "Read capabilities (default: true)",
				},
				cli.BoolFlag{
					Name:  "privacy, p",
					Usage: "Enable privacy mode (default: false)",
				},
			},
		})
}

// ActionMod provides the add subcommand
func ActionMod(c *cli.Context) {
	if len(c.Args()) != 1 {
		log.Fatal("error: must specify an app id")
	}
	app := c.Args()[0]
	_, err := service.Service.AppMod(app, &securepass.ApplicationDescriptor{
		Label:            c.String("label"),
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

	log.Println()
}
