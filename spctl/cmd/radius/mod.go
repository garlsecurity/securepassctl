package radius

import (
	"log"

	"github.com/urfave/cli"
	"github.com/garlsecurity/securepassctl"
	"github.com/garlsecurity/securepassctl/spctl/service"
)

func init() {
	Command.Subcommands = append(Command.Subcommands,
		cli.Command{
			Name:        "mod",
			Usage:       "modify RADIUS",
			ArgsUsage:   "RADIUS",
			Description: "Modify a RADIUS device in SecurePass.",
			Action:      ActionMod,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "realm, r",
					Usage: "Set alternate realm",
				},
				cli.StringFlag{
					Name:  "fqdn, f",
					Usage: "FQDN/Name",
				},
				cli.StringFlag{
					Name:  "secret, s",
					Usage: "Shared secret",
				},
				cli.StringFlag{
					Name:  "group, g",
					Usage: "Group name (restriction)",
				},
				cli.BoolFlag{
					Name:  "rfid, R",
					Usage: "Enable RFID tag as username",
				},
				cli.BoolTFlag{
					Name:  "no-rfid, n",
					Usage: "Disable RFID tag as username",
				},
			},
		})
}

// ActionMod provides the radius mod command
func ActionMod(c *cli.Context) {
	if len(c.Args()) != 1 {
		log.Fatal("error: must specify an IP address")
	}
	radius := c.Args()[0]
	securepassctl.DebugLogger.Printf("Modifying RADIUS %q", radius)
	_, err := service.Service.RadiusMod(radius, &securepassctl.RadiusDescriptor{
		Name:   c.String("fqdn"),
		Secret: c.String("secret"),
		Group:  c.String("group"),
		Realm:  c.String("realm"),
		Rfid:   c.Bool("rfid"),
	})
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
